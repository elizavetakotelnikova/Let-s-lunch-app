package usecases

import (
	"cmd/app/entities/user"
	domain "cmd/app/entities/user"
	"cmd/app/entities/user/query"
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
	PhoneNumber string
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
		command.PhoneNumber,
		command.Gender,
	)

	// Is there any user with same username
	criteria := query.FindCriteria{}
	criteria.Username.String = user.Username
	criteria.Username.Valid = true

	existingUser, err := useCase.user.FindUsersByCriteria(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("user: create user %w", err)
	}
	if len(existingUser) == 0 {
		_, err = useCase.user.Create(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("user: create user %w", err)
		}
		return user, nil
	} else {
		return nil, fmt.Errorf("user: username already exists")
	}
}
