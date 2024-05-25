package authbp

import (
	"crypto/rand"
	"encoding/base64"
)

type base64RefreshTokenProvider struct{}

func NewBase64RefreshTokenProvider() RefreshTokenProvider {
	return &base64RefreshTokenProvider{}
}

func (p *base64RefreshTokenProvider) Create() (string, error) {
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}

	token := base64.URLEncoding.EncodeToString(tokenBytes)

	return token, nil
}
