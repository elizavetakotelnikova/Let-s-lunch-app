package usecases

import (
	"cmd/app/entities/user"
	domain "cmd/app/entities/user"
	"cmd/app/entities/user/repository"
	"context"
	"fmt"
	"time"
)

type CreateUserUseCase struct {
	user repository.UsersRepository
}

type CreateUserCommand struct {
	Username    string
	DisplayName string
	Birthday    time.Time
	Gender      user.Gender
}

func NewCreateUserUseCase(user repository.UsersRepository) *CreateUserUseCase {
	return &CreateUserUseCase{user: user}
}

func (useCase *CreateUserUseCase) Handle(
	ctx context.Context,
	command CreateUserCommand,
) (*domain.User, error) {
	user := user.NewUser(
		command.Username,
		command.DisplayName,
		command.Birthday,
		command.Gender,
	)

	_, err := useCase.user.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("user: create user %w", err)
	}

	return user, nil
}
