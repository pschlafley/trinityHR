package api

import (
	"fmt"
	"net/http"

	"github.com/gofor-little/env"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/pschlafley/trinityHR/types"
)

func createJWT(account *types.Account) (string, error) {
	claims := &jwt.MapClaims{
		"ExpiresAt": 15000,
		"accountID": account.AccountID,
	}

	secret := env.Get("AuthSecret", "")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	var secret = env.Get("AuthSecret", "")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
}

func (s *APIServer) withJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-jwt-token")

		token, err := validateJWT(tokenString)

		if err != nil {
			permissionDenied(w)
			return
		}

		if !token.Valid {
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		var userReqID float64 = claims["accountID"].(float64)

		if userReqID == 0 {
			permissionDenied(w)
			return
		}

		account, err := s.store.GetAccountByJWT(token)

		if err != nil {
			permissionDenied(w)
			return
		}

		if userReqID != float64(account.AccountID) {
			permissionDenied(w)
			return
		}

		handlerFunc(w, r)

	}
}
