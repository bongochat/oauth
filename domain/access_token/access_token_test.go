package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstant(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration time must be 24 hours.")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "New access token should not expired")
	assert.EqualValues(t, at.AccessToken, "", "New access token should not empty")
	assert.EqualValues(t, at.PhoneNumber, "", "Phone number should not be empty")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "Access token should be expired by default")

	at.Expires = time.Now().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Expired access token created three thours from now should not be expired")
}
