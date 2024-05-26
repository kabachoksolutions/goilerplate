package authbp

import (
	"fmt"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type jwtTokenProvider struct {
	privateKey []byte
	issuer     string
}

func NewJWTTokenProvider(privateKey, issuer string) AccessTokenProvider {
	return &jwtTokenProvider{
		privateKey: []byte(privateKey),
		issuer:     issuer,
	}
}

func (p *jwtTokenProvider) Create(claims map[string]interface{}) (string, error) {
	if len(p.privateKey) == 0 {
		return "", fmt.Errorf("authbp: jwtTokenProvider: private key cannot be empty")
	}

	builder := jwt.NewBuilder().Issuer(p.issuer)

	for k, v := range claims {
		builder.Claim(k, v)
	}

	token, err := builder.IssuedAt(time.Now()).Build()
	if err != nil {
		return "", fmt.Errorf("authbp: jwtTokenProvider: failed to build token: %w", err)
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, p.privateKey))
	if err != nil {
		return "", fmt.Errorf("authbp: jwtTokenProvider: failed to sign token: %w", err)
	}

	return string(signed), nil
}

func (p *jwtTokenProvider) Verify(token []byte) (jwt.Token, error) {
	verifiedToken, err := jwt.Parse(token, jwt.WithKey(jwa.HS256, p.privateKey))
	if err != nil {
		return nil, fmt.Errorf("authbp: jwtTokenProvider: failed to verify token: %w", err)
	}

	return verifiedToken, nil
}
