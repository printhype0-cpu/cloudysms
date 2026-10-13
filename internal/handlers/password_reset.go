package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/smtp"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shridarpatil/whatomate/internal/models"
	"github.com/valyala/fasthttp"
	"github.com/zerodha/fastglue"
	"golang.org/x/crypto/bcrypt"
)

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

// generateSecureToken creates a random hex string for tokens
func generateSecureToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// getBaseURL extracts the client base URL from the request or falls back to port 3000
func getBaseURL(r *fastglue.Request) string {
	if origin := string(r.RequestCtx.Request.Header.Peek("Origin")); origin != "" {
		return strings.TrimRight(origin, "/")
	}
	if referer := string(r.RequestCtx.Request.Header.Peek("Referer")); referer != "" {
		if u, err := url.Parse(referer); err == nil && u.Host != "" {
			return fmt.Sprintf("%s://%s", u.Scheme, u.Host)
		}
	}
	host := string(r.RequestCtx.Host())
	if host != "" {
		scheme := "http"
		if string(r.RequestCtx.Request.Header.Peek("X-Forwarded-Proto")) == "https" || r.RequestCtx.IsTLS() {
			scheme = "https"
		}
		return fmt.Sprintf("%s://%s", scheme, host)
	}
	return "http://localhost:3000"
}

// ForgotPassword handles sending the password reset email/link
func (a *App) ForgotPassword(r *fastglue.Request) error {
	var req ForgotPasswordRequest
	if err := a.decodeRequest(r, &req); err != nil {
		return nil
	}

	req.Email = strings.TrimSpace(req.Email)
	if req.Email == "" {
		return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "Email is required", nil, "")
	}

	var user models.User
	if err := a.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		// Do not leak whether the email exists or not. Return success anyway.
		return r.SendEnvelope(map[string]interface{}{
			"message": "If that email address is in our database, we will send you an email to reset your password.",
		})
	}

	tokenStr, err := generateSecureToken()
	if err != nil {
		a.Log.Error("Failed to generate secure token", "error", err)
		return r.SendErrorEnvelope(fasthttp.StatusInternalServerError, "Internal error", nil, "")
	}

	// Invalidate any previously unused reset tokens for this user
	_ = a.DB.Model(&models.PasswordResetToken{}).
		Where("user_id = ? AND used = ?", user.ID, false).
		Update("used", true).Error

	resetToken := models.PasswordResetToken{
		BaseModel: models.BaseModel{ID: uuid.New()},
		UserID:    user.ID,
		Token:     tokenStr,
		ExpiresAt: time.Now().UTC().Add(24 * time.Hour),
		Used:      false,
	}

	if err := a.DB.Create(&resetToken).Error; err != nil {
		a.Log.Error("Failed to save reset token", "error", err)
		return r.SendErrorEnvelope(fasthttp.StatusInternalServerError, "Internal error", nil, "")
	}

	baseURL := getBaseURL(r)
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", baseURL, tokenStr)

	if a.Config.SMTP.Host != "" && a.Config.SMTP.Port > 0 {
		// Send email via SMTP
		auth := smtp.PlainAuth("", a.Config.SMTP.Username, a.Config.SMTP.Password, a.Config.SMTP.Host)
		to := []string{user.Email}
		from := a.Config.SMTP.FromEmail
		if from == "" {
			from = "noreply@whatomate.com"
		}

		subject := "Subject: Reset Your Password\r\n"
		mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

		body := fmt.Sprintf(`
			<p>Hello %s,</p>
			<p>We received a request to reset your password.</p>
			<p>Click the link below to reset it:</p>
			<p><a href="%s">Reset Password</a></p>
			<p>This link will expire in 24 hours.</p>
		`, user.FullName, resetLink)

		msg := []byte(subject + mime + body)

		addr := fmt.Sprintf("%s:%d", a.Config.SMTP.Host, a.Config.SMTP.Port)
		go func() {
			if err := smtp.SendMail(addr, auth, from, to, msg); err != nil {
				a.Log.Error("Failed to send password reset email", "error", err, "email", user.Email)
			}
		}()
	} else {
		// No SMTP configured, log it to standard output (development mode)
		a.Log.Info("========================================")
		a.Log.Info("PASSWORD RESET LINK GENERATED (NO SMTP CONFIG)")
		a.Log.Info("Email: " + user.Email)
		a.Log.Info("Link:  " + resetLink)
		a.Log.Info("========================================")
	}

	return r.SendEnvelope(map[string]interface{}{
		"message": "If that email address is in our database, we will send you an email to reset your password.",
	})
}

// VerifyResetToken validates whether a token is valid, unexpired, and unused
func (a *App) VerifyResetToken(r *fastglue.Request) error {
	tokenStr := strings.TrimSpace(string(r.RequestCtx.QueryArgs().Peek("token")))
	if tokenStr == "" {
		return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "Reset token is required", nil, "")
	}

	var token models.PasswordResetToken
	if err := a.DB.Where("token = ?", tokenStr).First(&token).Error; err != nil {
		return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "Invalid password reset link. Please request a new one.", nil, "")
	}

	if token.Used {
		return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "This password reset link has already been used.", nil, "")
	}

	if time.Now().UTC().After(token.ExpiresAt.UTC()) {
		return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "This password reset link has expired. Please request a new one.", nil, "")
	}

	return r.SendEnvelope(map[string]interface{}{
		"valid": true,
	})
}

// ResetPassword handles the actual password reset using a token
func (a *App) ResetPassword(r *fastglue.Request) error {
	var req ResetPasswordRequest
	if err := a.decodeRequest(r, &req); err != nil {
		return nil
	}

	req.Token = strings.TrimSpace(req.Token)
	if req.Token == "" || req.Password == "" {
		return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "Token and new password are required", nil, "")
	}

	if len(req.Password) < 8 {
		return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "Password must be at least 8 characters long", nil, "")
	}

	tx := a.DB.Begin()
	if tx.Error != nil {
		a.Log.Error("Failed to start transaction", "error", tx.Error)
		return r.SendErrorEnvelope(fasthttp.StatusInternalServerError, "Internal error", nil, "")
	}

	var token models.PasswordResetToken
	if err := tx.Where("token = ?", req.Token).First(&token).Error; err != nil {
		tx.Rollback()
		return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "Invalid password reset link. Please request a new one.", nil, "")
	}

	if token.Used {
		tx.Rollback()
		return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "This password reset link has already been used.", nil, "")
	}

	if time.Now().UTC().After(token.ExpiresAt.UTC()) {
		tx.Rollback()
		return r.SendErrorEnvelope(fasthttp.StatusBadRequest, "This password reset link has expired. Please request a new one.", nil, "")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		a.Log.Error("Failed to hash new password", "error", err)
		return r.SendErrorEnvelope(fasthttp.StatusInternalServerError, "Internal error", nil, "")
	}

	if err := tx.Model(&models.User{}).Where("id = ?", token.UserID).Update("password_hash", string(hashedPassword)).Error; err != nil {
		tx.Rollback()
		a.Log.Error("Failed to update user password", "error", err)
		return r.SendErrorEnvelope(fasthttp.StatusInternalServerError, "Internal error", nil, "")
	}

	if err := tx.Model(&token).Update("used", true).Error; err != nil {
		tx.Rollback()
		a.Log.Error("Failed to mark token as used", "error", err)
		return r.SendErrorEnvelope(fasthttp.StatusInternalServerError, "Internal error", nil, "")
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.Error("Failed to commit password reset transaction", "error", err)
		return r.SendErrorEnvelope(fasthttp.StatusInternalServerError, "Internal error", nil, "")
	}

	return r.SendEnvelope(map[string]interface{}{
		"message": "Password successfully reset. You can now log in.",
	})
}
