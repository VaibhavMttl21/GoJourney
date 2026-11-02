package user

import (
	"context"
	"errors"
	"sync"
)

type InMemoryRepository struct{
	mu sync.RWMutex
	users map[string]User
}

func NewInMemoryRepository () *InMemoryRepository{
	return &InMemoryRepository{
		users : make(map[string]User),
	}
}

func (r * InMemoryRepository) Create(ctx context.Context, user User) (User,error){
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, existinguser := range r.users{
		if(existinguser.Email==user.Email){
			return user,errors.New("Email already exist")
		}
	}

	r.users[user.ID] = user

	return user,nil
}

func (r *InMemoryRepository) FindbyID(ctx context.Context, id string)(User, error){
	r.mu.RLock()
	defer r.mu.RUnlock()

	user , exists := r.users[id]

	if(!exists){
		return User{},ErrUserNotFound
	}
	return user, nil
}

func (r * InMemoryRepository) FindbyEmail(ctx context.Context, email string)(User, error){
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return User{}, ErrUserNotFound
}