package DbTypes

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.step.sm/crypto/randutil"
)

type AuthenticationToken struct {
	Token string `json:"token"`
}

type MyCustomClaims struct {
	RegisteredClaims jwt.RegisteredClaims `json:"registered_claims"`
}

func (s *AuthenticationToken) NewAuthenticationToken() (string, error) {
	secret, err := randutil.ASCII(20)

	mySigningKey := []byte(string(secret))

	if err != nil {
		return "", err
	}

	claims := MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims.RegisteredClaims)

	myToken, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return myToken, nil
}

func ValidateToken(token jwt.Token) error {

	return nil
}
