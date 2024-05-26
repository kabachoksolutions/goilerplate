package authbp

import (
	"crypto/rand"
	"encoding/base64"
)

type cryptoRefreshTokenProvider struct{}

func NewCryptoRefreshTokenProvider() RefreshTokenProvider {
	return &cryptoRefreshTokenProvider{}
}

func (p *cryptoRefreshTokenProvider) Create() (string, error) {
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}

	token := base64.URLEncoding.EncodeToString(tokenBytes)

	return token, nil
}
