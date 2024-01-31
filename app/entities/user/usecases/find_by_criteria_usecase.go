package usecases

import "cmd/app/entities/user/repository"

type FindUserByCriteriaUseCase struct {
	user repository.UsersRepository
}
