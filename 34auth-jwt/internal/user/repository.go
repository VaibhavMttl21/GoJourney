package user

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type Repository interface{
	Create(ctx context.Context,user User) (User,error)
	FindbyID(ctx context.Context, id string)(User, error)
	FindbyEmail(ctx context.Context, email string)(User, error)
}