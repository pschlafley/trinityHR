package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofor-little/env"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/pschlafley/trinityHR/types"
)

type JwtCustomClaims struct {
	AccountID   int    `json:"account_id"`
	AccountType string `json:"account_type"`
	jwt.RegisteredClaims
}

var registeredClaims jwt.RegisteredClaims = jwt.RegisteredClaims{
	ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	IssuedAt:  jwt.NewNumericDate(time.Now()),
	NotBefore: jwt.NewNumericDate(time.Now()),
}

func createJWT(account *types.Account) (string, error) {

	accountId := account.AccountID
	accountType := account.AccountType

	claims := &JwtCustomClaims{
		accountId,
		accountType,
		registeredClaims,
	}

	secret := env.Get("AuthSecret", "")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

// func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
// 	cookie := new(http.Cookie)
// 	cookie.Name = name
// 	cookie.Value = token
// 	cookie.Expires = expiration
// 	cookie.Path = "/"
// 	cookie.HttpOnly = true

// 	c.SetCookie(cookie)
// }

func validateJWT(tokenString string) (*jwt.Token, error) {
	var secret = env.Get("AuthSecret", "")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func permissionDenied(c echo.Context) {
	c.JSON(http.StatusForbidden, ApiError{Error: "permission denied"})
}

func (s *APIServer) withJWTAuth(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("x-jwt-token")

		token, err := validateJWT(tokenString)

		if err != nil {
			permissionDenied(c)
			return err
		}

		if !token.Valid {
			permissionDenied(c)
			return err
		}

		claims := token.Claims.(jwt.MapClaims)

		var userReqID float64 = claims["account_id"].(float64)

		if userReqID == 0 {
			permissionDenied(c)
			return err
		}

		account, err := s.store.GetAccountByJWT(token)

		if err != nil {
			permissionDenied(c)
			return err
		}

		if userReqID != float64(account.AccountID) {
			permissionDenied(c)
			return err
		}

		return handlerFunc(c)

	}
}
