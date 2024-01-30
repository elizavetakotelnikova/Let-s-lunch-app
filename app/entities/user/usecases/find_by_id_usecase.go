package usecases

import (
	"context"

	domain "cmd/app/entities/user"
	"cmd/app/entities/user/repository"
	"github.com/gofrs/uuid/v5"
)

type FindUserByIdResponse struct {
	User *domain.User `json:"data"`
}

type FindUserByIdUseCase struct {
	users repository.UsersRepository
}

func NewFindUserByIdUseCase(users repository.UsersRepository) *FindUserByIdUseCase {
	return &FindUserByIdUseCase{users: users}
}

func (useCase *FindUserByIdUseCase) Handle(ctx context.Context, id uuid.UUID) (*FindUserByIdResponse, error) {
	users, err := useCase.users.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &FindUserByIdResponse{
		User: users,
	}, nil
}
