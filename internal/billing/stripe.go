package billing

import (
	"fmt"
	"log"

	"github.com/shridarpatil/whatomate/internal/config"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
	"github.com/stripe/stripe-go/v78/webhook"
)

type StripeService struct {
	config *config.Config
}

func NewStripeService(cfg *config.Config) *StripeService {
	stripe.Key = cfg.Stripe.SecretKey
	return &StripeService{
		config: cfg,
	}
}

// CreateCheckoutSession creates a new Stripe Checkout session for a specific organization
func (s *StripeService) CreateCheckoutSession(orgID string, successURL string, cancelURL string) (string, error) {
	if s.config.Stripe.SecretKey == "" || s.config.Stripe.PriceID == "" {
		return "", fmt.Errorf("stripe is not configured properly")
	}

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Mode:               stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(s.config.Stripe.PriceID),
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String(successURL),
		CancelURL:  stripe.String(cancelURL),
		ClientReferenceID: stripe.String(orgID), // Store Organization ID to map it back later
	}

	sess, err := session.New(params)
	if err != nil {
		log.Printf("Failed to create Stripe checkout session: %v", err)
		return "", err
	}

	return sess.URL, nil
}

// ConstructEvent parses the webhook payload securely
func (s *StripeService) ConstructEvent(payload []byte, signature string) (stripe.Event, error) {
	return webhook.ConstructEvent(payload, signature, s.config.Stripe.WebhookSecret)
}
