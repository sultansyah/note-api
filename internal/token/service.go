package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenService interface {
	GenerateToken(userId int, role string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type TokenServiceImpl struct {
	key string
}

func NewTokenService(key string) TokenService {
	return &TokenServiceImpl{key: key}
}

func (t *TokenServiceImpl) GenerateToken(userId int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 3).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(t.key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (t *TokenServiceImpl) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return t.key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, errors.New("token has expired")
			}
		}
		return token, nil
	}

	return nil, errors.New("invalid token claims")
}
