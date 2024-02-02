package usecases

import (
	"cmd/app/entities/user/query"
	"cmd/app/entities/user/repository"
	"context"
	"database/sql"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

type GetTokenUseCase struct {
	users     repository.UsersRepository
	tokenAuth *jwtauth.JWTAuth
}

func NewGetTokenUseCase(users repository.UsersRepository, tokenAuth *jwtauth.JWTAuth) *GetTokenUseCase {
	return &GetTokenUseCase{
		users:     users,
		tokenAuth: tokenAuth,
	}
}

func (t *GetTokenUseCase) Handle(ctx context.Context, username string, password string) (string, error) {
	users, err := t.users.FindUsersByCriteria(ctx, query.FindCriteria{
		Username: sql.NullString{String: username, Valid: true},
	})
	if err != nil {
		return "", err
	}
	if len(users) != 1 {
		return "", errors.New("user not found")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	if users[0].Username != username || string(users[0].HashedPassword) != string(hashed) {
		return "", errors.New("invalid username of password")
	}

	_, tokenStr, err := t.tokenAuth.Encode(map[string]interface{}{
		"username": username,
		"password": password,
	})
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
