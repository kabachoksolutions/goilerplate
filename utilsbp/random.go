package utilsbp

import (
	"crypto/rand"
	"math/big"
)

type RandomCodeType string

const (
	RandomCodeTypeOnlyNumbers       RandomCodeType = "only_numbers"
	RandomCodeTypeOnlyLetters       RandomCodeType = "only_letters"
	RandomCodeTypeLettersAndNumbers RandomCodeType = "letters_and_numbers"
)

func GenerateRandomCode(length int, codeType RandomCodeType) string {
	var chars string
	switch codeType {
	case RandomCodeTypeOnlyNumbers:
		chars = "0123456789"
	case RandomCodeTypeOnlyLetters:
		chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case RandomCodeTypeLettersAndNumbers:
		chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	default:
		return ""
	}

	code := make([]byte, length)
	max := big.NewInt(int64(len(chars)))

	for i := range code {
		randomIndex, err := rand.Int(rand.Reader, max)
		if err != nil {
			return ""
		}
		code[i] = chars[randomIndex.Int64()]
	}

	return string(code)
}
