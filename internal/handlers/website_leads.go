package handlers

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/valyala/fasthttp"
	"github.com/zerodha/fastglue"
)

type WebsiteLeadRequest struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

// HandleWebsiteLead captures leads from the landing page live chat
func (a *App) HandleWebsiteLead(r *fastglue.Request) error {
	var req WebsiteLeadRequest
	if err := a.decodeRequest(r, &req); err != nil {
		return nil
	}

	req.Email = strings.TrimSpace(req.Email)
	if req.Email == "" {
		return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "Email is required", nil, "")
	}

	adminEmail := a.Config.DefaultAdmin.Email
	if adminEmail == "" {
		adminEmail = "admin@admin.com"
	}

	if a.Config.SMTP.Host != "" && a.Config.SMTP.Port > 0 {
		// Send email via SMTP
		auth := smtp.PlainAuth("", a.Config.SMTP.Username, a.Config.SMTP.Password, a.Config.SMTP.Host)
		to := []string{adminEmail}
		from := a.Config.SMTP.FromEmail
		if from == "" {
			from = "noreply@whatomate.com"
		}

		subject := "Subject: New Website Lead from Live Chat\r\n"
		mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		
		body := fmt.Sprintf(`
			<h2>New Website Lead</h2>
			<p><strong>Visitor Email:</strong> %s</p>
			<p><strong>Message:</strong></p>
			<blockquote>%s</blockquote>
			<br/>
			<p>Please reply to them ASAP!</p>
		`, req.Email, req.Message)

		msg := []byte(subject + mime + body)

		addr := fmt.Sprintf("%s:%d", a.Config.SMTP.Host, a.Config.SMTP.Port)
		go func() {
			if err := smtp.SendMail(addr, auth, from, to, msg); err != nil {
				a.Log.Error("Failed to send website lead email", "error", err, "lead_email", req.Email)
			}
		}()
	} else {
		// No SMTP configured, log it to standard output
		a.Log.Info("========================================")
		a.Log.Info("NEW WEBSITE LEAD CAPTURED (NO SMTP CONFIG)")
		a.Log.Info("From: " + req.Email)
		a.Log.Info("Message: " + req.Message)
		a.Log.Info("========================================")
	}

	return r.SendEnvelope(map[string]interface{}{
		"success": true,
		"message": "Lead captured successfully",
	})
}
