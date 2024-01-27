package tests

import (
	"cmd/app/entities/gatheringPlace"
	repository2 "cmd/app/entities/gatheringPlace/repository"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "foodate"
)

func TestMeetingsRepositoryPassedCorrectShouldReturnSaved(t *testing.T) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	var databasePlacesRepository = repository2.NewPlacesDatabaseRepository(db)
	var currentPlace = gatheringPlace.NewGatheringPlace()
	currentPlace.CuisineType = gatheringPlace.FastFood
	var ctx = context.Background()
	_, err = databasePlacesRepository.Create(ctx, currentPlace)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
