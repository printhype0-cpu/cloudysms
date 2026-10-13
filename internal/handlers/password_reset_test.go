package handlers_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shridarpatil/whatomate/internal/models"
	"github.com/shridarpatil/whatomate/test/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
)

func TestApp_ForgotPassword_SuccessAndNoPKConflict(t *testing.T) {
	app := newTestApp(t)
	org := testutil.CreateTestOrganization(t, app.DB)
	email := testutil.UniqueEmail("forgot-pwd")
	user := testutil.CreateTestUser(t, app.DB, org.ID, testutil.WithEmail(email))

	// First request: should create reset token
	req1 := testutil.NewJSONRequest(t, map[string]string{
		"email": email,
	})
	req1.RequestCtx.Request.Header.Set("Origin", "http://localhost:3000")

	err := app.ForgotPassword(req1)
	require.NoError(t, err)
	assert.Equal(t, fasthttp.StatusOK, testutil.GetResponseStatusCode(req1))

	var tokens1 []models.PasswordResetToken
	err = app.DB.Where("user_id = ?", user.ID).Find(&tokens1).Error
	require.NoError(t, err)
	require.Len(t, tokens1, 1)
	assert.NotEqual(t, uuid.Nil, tokens1[0].ID)
	assert.False(t, tokens1[0].Used)
	assert.True(t, tokens1[0].ExpiresAt.After(time.Now().UTC().Add(23*time.Hour)))

	// Second request: tests that subsequent tokens do NOT collide on primary key (UUID)
	// and that previous tokens are invalidated
	req2 := testutil.NewJSONRequest(t, map[string]string{
		"email": email,
	})
	err = app.ForgotPassword(req2)
	require.NoError(t, err)
	assert.Equal(t, fasthttp.StatusOK, testutil.GetResponseStatusCode(req2))

	var tokens2 []models.PasswordResetToken
	err = app.DB.Where("user_id = ?", user.ID).Order("created_at ASC").Find(&tokens2).Error
	require.NoError(t, err)
	require.Len(t, tokens2, 2)
	assert.NotEqual(t, uuid.Nil, tokens2[1].ID)
	assert.NotEqual(t, tokens2[0].ID, tokens2[1].ID)
	assert.True(t, tokens2[0].Used, "Previous token should be invalidated")
	assert.False(t, tokens2[1].Used, "New token should be active")
}

func TestApp_VerifyResetToken(t *testing.T) {
	app := newTestApp(t)
	org := testutil.CreateTestOrganization(t, app.DB)
	user := testutil.CreateTestUser(t, app.DB, org.ID)

	// Valid token
	validToken := models.PasswordResetToken{
		BaseModel: models.BaseModel{ID: uuid.New()},
		UserID:    user.ID,
		Token:     "valid-test-token",
		ExpiresAt: time.Now().UTC().Add(24 * time.Hour),
		Used:      false,
	}
	require.NoError(t, app.DB.Create(&validToken).Error)

	// Expired token
	expiredToken := models.PasswordResetToken{
		BaseModel: models.BaseModel{ID: uuid.New()},
		UserID:    user.ID,
		Token:     "expired-test-token",
		ExpiresAt: time.Now().UTC().Add(-1 * time.Hour),
		Used:      false,
	}
	require.NoError(t, app.DB.Create(&expiredToken).Error)

	// Used token
	usedToken := models.PasswordResetToken{
		BaseModel: models.BaseModel{ID: uuid.New()},
		UserID:    user.ID,
		Token:     "used-test-token",
		ExpiresAt: time.Now().UTC().Add(24 * time.Hour),
		Used:      true,
	}
	require.NoError(t, app.DB.Create(&usedToken).Error)

	// Test valid token
	reqValid := testutil.NewRequest(t)
	reqValid.RequestCtx.URI().SetQueryString("token=valid-test-token")
	err := app.VerifyResetToken(reqValid)
	require.NoError(t, err)
	assert.Equal(t, fasthttp.StatusOK, testutil.GetResponseStatusCode(reqValid))

	// Test expired token
	reqExpired := testutil.NewRequest(t)
	reqExpired.RequestCtx.URI().SetQueryString("token=expired-test-token")
	err = app.VerifyResetToken(reqExpired)
	require.NoError(t, err)
	assert.Equal(t, fasthttp.StatusBadRequest, testutil.GetResponseStatusCode(reqExpired))
	var expiredResp struct {
		Message string `json:"message"`
	}
	_ = json.Unmarshal(testutil.GetResponseBody(reqExpired), &expiredResp)
	assert.Contains(t, expiredResp.Message, "expired")

	// Test used token
	reqUsed := testutil.NewRequest(t)
	reqUsed.RequestCtx.URI().SetQueryString("token=used-test-token")
	err = app.VerifyResetToken(reqUsed)
	require.NoError(t, err)
	assert.Equal(t, fasthttp.StatusBadRequest, testutil.GetResponseStatusCode(reqUsed))
	var usedResp struct {
		Message string `json:"message"`
	}
	_ = json.Unmarshal(testutil.GetResponseBody(reqUsed), &usedResp)
	assert.Contains(t, usedResp.Message, "already been used")

	// Test non-existent token
	reqNotFound := testutil.NewRequest(t)
	reqNotFound.RequestCtx.URI().SetQueryString("token=nonexistent-token")
	err = app.VerifyResetToken(reqNotFound)
	require.NoError(t, err)
	assert.Equal(t, fasthttp.StatusBadRequest, testutil.GetResponseStatusCode(reqNotFound))
}

func TestApp_ResetPassword_Success(t *testing.T) {
	app := newTestApp(t)
	org := testutil.CreateTestOrganization(t, app.DB)
	user := testutil.CreateTestUser(t, app.DB, org.ID)

	tokenStr := "valid-reset-token-12345"
	token := models.PasswordResetToken{
		BaseModel: models.BaseModel{ID: uuid.New()},
		UserID:    user.ID,
		Token:     tokenStr,
		ExpiresAt: time.Now().UTC().Add(24 * time.Hour),
		Used:      false,
	}
	require.NoError(t, app.DB.Create(&token).Error)

	newPassword := "NewSecurePassword123"
	req := testutil.NewJSONRequest(t, map[string]string{
		"token":    tokenStr,
		"password": newPassword,
	})

	err := app.ResetPassword(req)
	require.NoError(t, err)
	assert.Equal(t, fasthttp.StatusOK, testutil.GetResponseStatusCode(req))

	// Verify token is now marked as used
	var updatedToken models.PasswordResetToken
	err = app.DB.Where("id = ?", token.ID).First(&updatedToken).Error
	require.NoError(t, err)
	assert.True(t, updatedToken.Used)

	// Verify user password hash was updated
	var updatedUser models.User
	err = app.DB.Where("id = ?", user.ID).First(&updatedUser).Error
	require.NoError(t, err)
	err = bcrypt.CompareHashAndPassword([]byte(updatedUser.PasswordHash), []byte(newPassword))
	assert.NoError(t, err, "User password hash should match new password")

	// Resetting again with the same token should fail with already used message
	reqSecond := testutil.NewJSONRequest(t, map[string]string{
		"token":    tokenStr,
		"password": "AnotherPassword123",
	})
	err = app.ResetPassword(reqSecond)
	require.NoError(t, err)
	assert.Equal(t, fasthttp.StatusBadRequest, testutil.GetResponseStatusCode(reqSecond))
	var secondResp struct {
		Message string `json:"message"`
	}
	_ = json.Unmarshal(testutil.GetResponseBody(reqSecond), &secondResp)
	assert.Contains(t, secondResp.Message, "already been used")
}
