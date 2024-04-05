package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateAccessTokenSuccess(t *testing.T) {
	_, err := GenerateAccessToken("TestAccess", "TestAccess")

	assert.Equal(t, err, nil)
}
func TestGenerateRefreshTokenSuccess(t *testing.T) {
	_, err := GenerateRefreshToken("TestRefresh", "TestRefresh")

	assert.Equal(t, err, nil)
}

func TestValidateTokenSuccess(t *testing.T) {
	accessToken, _ := GenerateAccessToken("TestAccess", "TestAccess")
	refreshToken, _ := GenerateRefreshToken("TestRefresh", "TestRefresh")

	err := ValidateToken(accessToken, "TestAccess")
	assert.Equal(t, err, nil)
	err = ValidateToken(refreshToken, "TestRefresh")
	assert.Equal(t, err, nil)
}
func TestValidateTokenExpiredTime(t *testing.T) {
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlRlc3RBY2Nlc3MiLCJleHAiOjE2ODM3MjMyNTN9.reVoYH1Z4u9AouONwqWeBEc2t5y9IxtxusESXnjra_E"
	refreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlRlc3RSZWZyZXNoIiwiZXhwIjoxNjgzNzM3NjUzfQ.6eNN3L14NJGBfNKdL6HbtrB4QCZD7U2PWkQqnftpbAM"

	err := ValidateToken(accessToken, "TestAccess")
	assert.NotEqual(t, err, nil)

	err = ValidateToken(refreshToken, "TestRefresh")
	assert.NotEqual(t, err, nil)
}
func TestExtractUsernameFromToken(t *testing.T) {
	accessToken, _ := GenerateAccessToken("TestAccess", "TestAccess")
	refreshToken, _ := GenerateRefreshToken("TestRefresh", "TestRefresh")

	accessUsername, err := ExtractUsernameFromToken(accessToken, "TestAccess")
	assert.Equal(t, accessUsername, "TestAccess")
	assert.Equal(t, err, nil)
	refreshUsername, err := ExtractUsernameFromToken(refreshToken, "TestRefresh")
	assert.Equal(t, refreshUsername, "TestRefresh")
	assert.Equal(t, err, nil)
}

//Invalid tokenlar için zaman lazım
//Access Token Invalid
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlRlc3RBY2Nlc3MiLCJleHAiOjE2ODM3MjMyNTN9.reVoYH1Z4u9AouONwqWeBEc2t5y9IxtxusESXnjra_E
//Refresh Token Invalid
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlRlc3RSZWZyZXNoIiwiZXhwIjoxNjgzNzM3NjUzfQ.6eNN3L14NJGBfNKdL6HbtrB4QCZD7U2PWkQqnftpbAM
