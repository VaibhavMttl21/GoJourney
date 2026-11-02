package auth

import (
	"auth-jwt/internal/user"
	"context"
	"errors"
	"time"
	"uuid"
)


var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Service struct{
	userRepo user.Repository
	refreshTokenStore *RefreshTokenStore
}

func NewService(
    userRepo user.Repository,
    refreshTokenStore *RefreshTokenStore,
) *Service {
    return &Service{
        userRepo:          userRepo,
        refreshTokenStore: refreshTokenStore,
    }
}

func (s * Service) Signup(ctx context.Context,email string, password string)(user.User,error){
	_,err := s.userRepo.FindbyEmail(ctx,email)

	if err == nil{
		return user.User{},ErrEmailAlreadyExists
	}
	if !errors.Is(err, user.ErrUserNotFound){
		return user.User{},err
	}

	passwordHash,err := HashPassword(password)

	if(err!=nil){
		return user.User{},err
	}

	NewUser := user.User{
		ID : uuid.New().String(),
		Email: email,
		PasswordHash: passwordHash,
		Role: "user",
		// Role:role,
	}

	return s.userRepo.Create(ctx,NewUser)
}



func (s *Service) Signin(
    ctx context.Context,
    email string,
    password string,
) (string, string, error) {

    foundUser, err := s.userRepo.FindbyEmail(ctx, email)

    if err != nil {
        if errors.Is(err, user.ErrUserNotFound) {
            return "", "", ErrInvalidCredentials
        }

        return "", "", err
    }

    err = CheckPassword(
        password,
        foundUser.PasswordHash,
    )

    if err != nil {
        return "", "", ErrInvalidCredentials
    }

    accessToken, err := GenerateAccessToken(
        foundUser.ID,
        foundUser.Role,
    )

    if err != nil {
        return "", "", err
    }

    refreshToken, err := GenerateRefreshToken()

    if err != nil {
        return "", "", err
    }

    s.refreshTokenStore.Save(
        refreshToken,
        foundUser.ID,
        time.Now().Add(7*24*time.Hour),
    )

    return accessToken, refreshToken, nil
}

