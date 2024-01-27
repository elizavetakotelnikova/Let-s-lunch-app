package tests

import (
	"cmd/app/entities/gatheringPlace"
	repository2 "cmd/app/entities/gatheringPlace/repository"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "foodate"
)

func TestCreatingGatheringPlace(t *testing.T) {
	//set up
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
	//main part
	_, errCreating := databasePlacesRepository.Create(ctx, currentPlace)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	placeFromRepository, errFinding := databasePlacesRepository.FindByID(ctx, currentPlace.ID)
	if errFinding != nil {
		t.Fatalf("Error in finding place: %s", errFinding)
	}
	//testing
	assert.Equal(t, currentPlace.ID, placeFromRepository.ID)
	assert.Equal(t, currentPlace.Address, placeFromRepository.Address)
	assert.Equal(t, currentPlace.AveragePrice, placeFromRepository.AveragePrice)
	assert.Equal(t, currentPlace.CuisineType, placeFromRepository.CuisineType)
	assert.Equal(t, currentPlace.PhoneNumber, placeFromRepository.PhoneNumber)
	err = databasePlacesRepository.Delete(ctx, currentPlace)
	if err != nil {
		t.Fatalf("Error in deleting place: %s", errFinding)
	}
}
