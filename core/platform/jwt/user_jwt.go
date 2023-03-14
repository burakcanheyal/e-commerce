package jwt

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateAccessToken(userName string, secret string) (string, error) {
	claims := &dto.UserToken{
		Username: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 3600)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func GenerateRefreshToken(userName string, secret2 string) (string, error) {
	claims := &dto.UserToken{
		Username: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 18000)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret2))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func ValidateToken(signedToken string, secret string) (err error) {

	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if token.Valid {
		return nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return jwt.ErrTokenMalformed
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return jwt.ErrTokenExpired
	} else {
		return internal.FailInToken
	}
}
func ExtractUsernameFromToken(requestToken string, secret string) (string, error) {

	token, err := jwt.ParseWithClaims(requestToken, &dto.UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*dto.UserToken); ok && token.Valid {
		return claims.Username, nil
	} else {
		return "", internal.FailInTokenParse
	}
}
