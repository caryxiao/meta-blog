package utils

import (
	"errors"
	"time"

	"github.com/caryxiao/meta-blog/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

// JWT configuration
var jwtConfig *config.JWT

// InitJWT initializes JWT configuration
func InitJWT(cfg *config.JWT) {
	jwtConfig = cfg
}

// Claims JWT payload
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken generates JWT token
func GenerateToken(userID uint, username string) (string, error) {
	if jwtConfig == nil {
		return "", errors.New("JWT config not initialized")
	}

	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(jwtConfig.ExpireHours) * time.Hour)

	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			Issuer:    "meta-blog",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.Secret))
}

// ParseToken parses JWT token
func ParseToken(tokenString string) (*Claims, error) {
	if jwtConfig == nil {
		return nil, errors.New("JWT config not initialized")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
