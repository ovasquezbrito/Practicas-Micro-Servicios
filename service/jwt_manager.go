package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTManager is a JSON web token manager
type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

type UserClains struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

// NewJWTManager returns a new JWTManager
func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenDuration}
}

//Generate generates and signs a new token for a user
func (manager *JWTManager) Generate(user *User) (string, error) {
	clains := UserClains{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		Username: user.Username,
		Role:     user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clains)
	return token.SignedString([]byte(manager.secretKey))
}

//Verify verifies the access token string and return a user clain if the token is valid
func (manager *JWTManager) Verify(accessToken string) (*UserClains, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClains{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	clains, ok := token.Claims.(*UserClains)
	if !ok {
		return nil, fmt.Errorf("invalid token clains")
	}

	return clains, nil

}
