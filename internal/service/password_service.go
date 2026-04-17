package service

import (
	"crypto/rand"
	"math/big"
)

type PasswordService interface {
	Generate(length int, includeUpper, includeLower, includeNumber, includeSymbol bool) (string, error)
}

type passwordService struct{}

func NewPasswordService() PasswordService {
	return &passwordService{}
}

func (s *passwordService) Generate(length int, includeUpper, includeLower, includeNumber, includeSymbol bool) (string, error) {
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower := "abcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	symbols := "!@#$%^&*()-_=+[]{}|;:,.<>?"

	charSet := ""
	if includeUpper {
		charSet += upper
	}
	if includeLower {
		charSet += lower
	}
	if includeNumber {
		charSet += numbers
	}
	if includeSymbol {
		charSet += symbols
	}

	// 如果没有选择任何字符集，默认使用数字和字母
	if charSet == "" {
		charSet = lower + numbers
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return "", err
		}
		password[i] = charSet[num.Int64()]
	}

	return string(password), nil
}
