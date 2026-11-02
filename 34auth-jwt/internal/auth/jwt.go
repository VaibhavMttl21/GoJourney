package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	// "google.golang.org/grpc/credentials/jwt"
)

var jwtSecret = []byte ("my-super-secret-key")

type Claims struct {
	// UserID string `json:"user_id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}
// So we've now put both identity and authorization information into the token.
func GenerateAccessToken(userID string, role string) (string, error) {

	now := time.Now()

	claims := Claims{
		Role: role,

		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(
				now.Add(15 * time.Minute),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken() (string, error) {
    bytes := make([]byte, 32)

    _, err := rand.Read(bytes)
    if err != nil {
        return "", err
    }

    return base64.RawURLEncoding.EncodeToString(bytes), nil
}

func (s *RefreshTokenStore) Get(
    token string,
) (RefreshToken, bool) {

    s.mu.RLock()
    defer s.mu.RUnlock()

    refreshToken, exists := s.tokens[token]

    if !exists {
        return RefreshToken{}, false
    }

    if time.Now().After(refreshToken.ExpiresAt) {
        return RefreshToken{}, false
    }

    return refreshToken, true
}

func ValidateAccessToken(tokenString string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {

			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New(
					"unexpected signing method",
				)
			}

			return jwtSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}