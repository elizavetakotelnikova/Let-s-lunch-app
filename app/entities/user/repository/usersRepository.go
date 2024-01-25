package repository

import (
	"cmd/app/entities/user"
	"cmd/app/entities/user/query"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

//go:generate mockery --name=UsersRepository
type UsersRepository interface {
	FindUsersByCriteria(ctx context.Context, criteria query.FindCriteria) ([]user.User, error)
	Create(ctx context.Context, user user.User) (*user.User, error)
	Update(ctx context.Context, user user.User) (*user.User, error)
	Delete(ctx context.Context, user user.User) error
}

type UsersDatabaseRepository struct {
	db *sql.DB
}

func NewUsersDatabaseRepository(providedConnection *sql.DB) *UsersDatabaseRepository {
	return &UsersDatabaseRepository{db: providedConnection}
}
func (repository *UsersDatabaseRepository) FindUsersByCriteria(ctx context.Context, criteria query.FindCriteria) ([]user.User, error) {
	var users []user.User
	rows, err := query.FindUserByCriteria(ctx, criteria, repository.db)
	if err != nil {
		return nil, fmt.Errorf("cannot query the database %w", err)
	}
	var currentUser user.User
	for rows.Next() {
		if err = rows.Scan(&currentUser.ID, &currentUser.Username, &currentUser.DisplayName, &currentUser.CurrentMeetingId, &currentUser.Rating); err != nil {
			return nil, fmt.Errorf("cannot query the database %w", err)
		}
		users = append(users, currentUser)
		rows, err = query.FindUserHistoryById(ctx, currentUser.ID, repository.db)
		if err != nil {
			return nil, fmt.Errorf("cannot query meeting history %w", err)
		}
		var meetingId uuid.UUID
		for rows.Next() {
			if err := rows.Scan(&meetingId); err != nil {
				return nil, fmt.Errorf("FindById %d: %v", currentUser.ID, err)
			}
			currentUser.MeetingHistory = append(currentUser.MeetingHistory, meetingId)
		}
	}
	return users, nil
}

func (repository *UsersDatabaseRepository) Create(ctx context.Context, user *user.User) (*user.User, error) {
	var err = query.Create(ctx, user, repository.db)
	if err != nil {
		return user, fmt.Errorf("meeting cannot be created: %v", err)
	}
	return nil, nil
}

func (repository *UsersDatabaseRepository) Update(ctx context.Context, user *user.User) (*user.User, error) {
	var err = query.Update(ctx, user, repository.db)
	if err != nil {
		return user, fmt.Errorf("meeting cannot be updated: %v", err)
	}
	return nil, nil
}

func (repository *UsersDatabaseRepository) Delete(ctx context.Context, user *user.User) error {
	var err = query.Delete(ctx, user, repository.db)
	if err != nil {

		return fmt.Errorf("meeting cannot be deleted: %v", err)
	}
	return nil
}
