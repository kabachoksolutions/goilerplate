package authbp

import "github.com/lestrrat-go/jwx/v2/jwt"

type PasswordProvider interface {
	Generate(password string) (string, error)
	Compare(password, encodedHash string) (bool, error)
}

type AccessTokenProvider interface {
	Create(claims map[string]interface{}) (string, error)
	Verify(token []byte) (jwt.Token, error)
}

type RefreshTokenProvider interface {
	Create() (string, error)
}
