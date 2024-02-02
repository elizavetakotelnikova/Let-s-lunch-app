package API_tests

import (
	"cmd/app/config"
	"cmd/di"
	"cmd/tests/mocks"
	"context"
	"database/sql"
	"fmt"
	"github.com/ory/dockertest/v3"
	_ "github.com/stretchr/testify"
	"log"
	"net/http"
	"os"
)

var server *http.Server

var usersRepository = mocks.UsersRepository{}
var gatheringPlacesRepository = mocks.PlacesRepository{}
var meetingsRepository = mocks.MeetingsRepository{}

func SetUpServer() {
	container, err := di.NewContainer(
		config.Params{
			DatabaseURL:   os.Getenv("DATABASE_URL"),
			ServerAddress: os.Getenv("SERVER_ADDRESS"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	server, err := container.Server(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	container.Close()
}

func setUpConnectionToTestContainers() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("postgres", "latest", []string{
		"POSTGRES_PASSWORD=secret",
		"POSTGRES_USER=user",
		"POSTGRES_DB=testdb",
	})

	var hostPort = resource.GetHostPort("5462/tcp")
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgresql", fmt.Sprintf("host=localhost port=%s user=user password=secret dbname=testdb sslmode=disable", hostPort))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

var db *sql.DB

func setUpConnection() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "admin"
		password = "admin"
		dbname   = "foodate"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
}

func closeConnection() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
