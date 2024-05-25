package authbp

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type bcryptPasswordProvider struct {
	cost int // cost factor for the bcrypt algorithm.
}

func NewBcryptPasswordProvider(cost int) PasswordProvider {
	if cost <= 0 {
		cost = bcrypt.DefaultCost
	}

	return &bcryptPasswordProvider{
		cost: cost,
	}
}

func (p *bcryptPasswordProvider) Generate(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), p.cost)
	if err != nil {
		return "", fmt.Errorf("bcryptPasswordProvider: failed to generate hash: %w", err)
	}

	return string(hashedPassword), nil
}

func (p *bcryptPasswordProvider) Compare(password, encodedHash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(encodedHash), []byte(password))
	if err != nil {
		return false, fmt.Errorf("bcryptPasswordProvider: failed to compare hash and password: %w", err)
	}

	return true, nil
}
