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
	Secret           string               `json:"secret"`
	RegisteredClaims jwt.RegisteredClaims `json:"registered_claims"`
}

func (s *AuthenticationToken) NewAuthenticationToken() (string, error) {
	secret, err := randutil.ASCII(25)

	mySigningKey := []byte("ALL YOUR BASE")

	if err != nil {
		return "", err
	}

	claims := MyCustomClaims{
		Secret: secret,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims.RegisteredClaims)

	ss, ssErr := token.SignedString(mySigningKey)

	if ssErr != nil {
		return "", ssErr
	}

	return ss, nil
}
