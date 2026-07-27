package handlers

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v78"
	"github.com/valyala/fasthttp"
	"github.com/zerodha/fastglue"
)

// CreateCheckoutSessionRequest represents the payload for creating a session
type CreateCheckoutSessionRequest struct {
	SuccessURL string `json:"success_url" validate:"required"`
	CancelURL  string `json:"cancel_url" validate:"required"`
}

// HandleCreateCheckoutSession creates a new Stripe checkout session
func (a *App) HandleCreateCheckoutSession(r *fastglue.Request) error {
	orgID, ok := r.RequestCtx.UserValue("organization_id").(uuid.UUID)
	if !ok {
		return r.SendErrorEnvelope(fasthttp.StatusUnauthorized, "Unauthorized", nil, "")
	}

	var req CreateCheckoutSessionRequest
	if err := a.decodeRequest(r, &req); err != nil {
		return nil
	}

	url, err := a.Billing.CreateCheckoutSession(orgID.String(), req.SuccessURL, req.CancelURL)
	if err != nil {
		a.Log.Error("Failed to create checkout session", "error", err)
		return r.SendErrorEnvelope(fasthttp.StatusInternalServerError, "Failed to create checkout session", nil, "")
	}

	return r.SendEnvelope(map[string]string{
		"url": url,
	})
}

// HandleStripeWebhook handles incoming webhooks from Stripe securely
func (a *App) HandleStripeWebhook(r *fastglue.Request) error {
	payload := r.RequestCtx.PostBody()
	signature := string(r.RequestCtx.Request.Header.Peek("Stripe-Signature"))

	event, err := a.Billing.ConstructEvent(payload, signature)
	if err != nil {
		a.Log.Error("Failed to verify Stripe webhook signature", "error", err)
		return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "Invalid webhook signature", nil, "")
	}

	switch event.Type {
	case "checkout.session.completed":
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			a.Log.Error("Failed to parse checkout session", "error", err)
			return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "Invalid payload", nil, "")
		}
		
		orgIDStr := session.ClientReferenceID
		if orgIDStr == "" {
			a.Log.Warn("Checkout session completed without ClientReferenceID")
			break
		}

		// Update organization
		if err := a.DB.Exec(`
			UPDATE organizations 
			SET stripe_customer_id = ?, 
			    stripe_subscription_id = ?, 
			    plan_tier = 'pro', 
			    subscription_status = 'active'
			WHERE id = ?`, 
			session.Customer.ID, 
			session.Subscription.ID, 
			orgIDStr,
		).Error; err != nil {
			a.Log.Error("Failed to update organization billing details", "error", err)
			return r.SendErrorEnvelope(fasthttp.StatusInternalServerError, "Failed to update database", nil, "")
		}
		
		a.Log.Info("Successfully processed checkout session", "orgID", orgIDStr)

	case "customer.subscription.updated", "customer.subscription.deleted":
		var subscription stripe.Subscription
		if err := json.Unmarshal(event.Data.Raw, &subscription); err != nil {
			a.Log.Error("Failed to parse subscription", "error", err)
			return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "Invalid payload", nil, "")
		}

		if err := a.DB.Exec(`
			UPDATE organizations 
			SET subscription_status = ?
			WHERE stripe_subscription_id = ?`, 
			string(subscription.Status), 
			subscription.ID,
		).Error; err != nil {
			a.Log.Error("Failed to update subscription status", "error", err)
			return r.SendErrorEnvelope(fasthttp.StatusInternalServerError, "Failed to update database", nil, "")
		}
		
		a.Log.Info("Successfully updated subscription status", "subID", subscription.ID, "status", subscription.Status)
	}

	r.RequestCtx.SetStatusCode(fasthttp.StatusOK)
	return nil
}
