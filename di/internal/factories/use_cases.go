package factories

import (
	gathering_place_usecase "cmd/app/entities/gatheringPlace/usecases"
	meeting_usecase "cmd/app/entities/meeting/usecases"
	user_usecase "cmd/app/entities/user/usecases"
	"cmd/di/internal/lookup"
	"context"
)

func CreateUseCasesFindMeeting(ctx context.Context, c lookup.Container) *meeting_usecase.FindMeetingByIdUseCase {
	return meeting_usecase.NewFindMeetingByIdUseCase(
		c.Repositories().MeetingRepository(ctx),
	)
}

func CreateUseCasesFindUser(ctx context.Context, c lookup.Container) *user_usecase.FindUserByIdUseCase {
	return user_usecase.NewFindUserByIdUseCase(
		c.Repositories().UserRepository(ctx),
	)
}

func CreateUseCasesFindGatheringPlace(ctx context.Context, c lookup.Container) *gathering_place_usecase.FindGatheringPlaceByIdUseCase {
	return gathering_place_usecase.NewFindGatheringPlaceByIdUseCase(
		c.Repositories().GatheringPlaceRepository(ctx),
	)
}

func CreateUseCasesCreateUser(ctx context.Context, c lookup.Container) *user_usecase.CreateUserUseCase {
	return user_usecase.NewCreateUserUseCase(
		c.Repositories().UserRepository(ctx),
	)
}
