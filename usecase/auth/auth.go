package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	Uid    int   `json:"uid"`
	Iat    int64 `json:"iat"`
	Expiry int64 `json:"expiry"`
}

func GenerateToken(userID string, now time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":    userID,
		"iat":    now.Unix(),
		"expiry": now.Add(12 * 24 * time.Hour),
	})

	return token.SignedString([]byte(os.Getenv("TOKEN_KRY")))
}

func ValidToken(signedString string) (*jwt.Token, error) {
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_KRY")), nil
	})

	if err != nil {
		return nil, error(err)
	}

	return token, nil
}

func ParseToken(signedString string) (*Auth, error) {
	token, err := ValidToken(signedString)

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("%s is expired", signedString)
			} else {
				return nil, fmt.Errorf("%s is invalid", signedString)
			}
		} else {
			return nil, fmt.Errorf("%s is invalid", signedString)
		}
	}

	if token == nil {
		return nil, fmt.Errorf("not found token in %s:", signedString)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("not found claims in %s", signedString)
	}

	userID, ok := claims["uid"].(int)
	if !ok {
		return nil, fmt.Errorf("not found uid in %s", signedString)
	}
	iat, ok := claims["iat"].(float64)
	if !ok {
		return nil, fmt.Errorf("not found iat in %s", signedString)
	}
	expiry, ok := claims["expiry"].(float64)
	if !ok {
		return nil, fmt.Errorf("not found expiry in %s", signedString)
	}

	return &Auth{
		Uid:    userID,
		Iat:    int64(iat),
		Expiry: int64(expiry),
	}, nil
}
