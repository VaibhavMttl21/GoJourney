package auth

import (
    "sync"
    "time"
)

type RefreshToken struct {
    Token     string
    UserID    string
    ExpiresAt time.Time
}

type RefreshTokenStore struct {
    mu     sync.RWMutex
    tokens map[string]RefreshToken
}

func NewRefreshTokenStore() *RefreshTokenStore {
    return &RefreshTokenStore{
        tokens: make(map[string]RefreshToken),
    }
}

func (s *RefreshTokenStore) Save(
    token string,
    userID string,
    expiresAt time.Time,
) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.tokens[token] = RefreshToken{
        Token:     token,
        UserID:    userID,
        ExpiresAt: expiresAt,
    }
}

func (s *RefreshTokenStore) Revoke(token string) {
    s.mu.Lock()
    defer s.mu.Unlock()

    delete(s.tokens, token)
}