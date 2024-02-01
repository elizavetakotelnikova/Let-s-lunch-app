package API_tests

import (
	"cmd/app/entities/user"
	"cmd/app/entities/user/api"
	"cmd/app/entities/user/repository"
	"cmd/app/entities/user/usecases"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"github.com/gorilla/mux"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"slices"
	"strings"
	"testing"
	"time"
)

func AreUsersEqual(first *user.User, second *user.User) bool {
	if (first.ID != second.ID) || (first.CurrentMeetingId != second.CurrentMeetingId) ||
		(!time.Time.Equal(first.Birthday, second.Birthday)) || (first.PhoneNumber != second.PhoneNumber) ||
		(first.Gender != second.Gender) || (!slices.Equal(first.MeetingHistory, second.MeetingHistory)) ||
		(first.Rating != second.Rating) || (first.Username != second.Username) || (first.DisplayName != second.DisplayName) {
		return false
	}
	return true
}

func TestMain(m *testing.M) {
	//setUpConnectionToTestContainers()
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
	opts := dockertest.RunOptions{
		Repository:   "postgres",
		Tag:          "14.8-alpine",
		Env:          []string{"POSTGRES_USER=admin POSTGRES_PASSWORD=admin DBNAME=foodate"},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5433": {
				{HostIP: "0.0.0.0", HostPort: "5432"},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf("host=localhost port=5432 user=admin password=admin dbname=foodate sslmode=disable"))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	code := m.Run()
	// You can't defer this because os.Exit doesn't care for defer
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

func TestCreatingUserShouldReturnStatus200(t *testing.T) {
	var requestBody = `{"username": "katya78", "displayName": "katya", "birthday": "2006-12-13T08:08:08Z", "phoneNumber": "89816999983", "gender": 1}`
	req := httptest.NewRequest(http.MethodPost, "/api/user/create", strings.NewReader(requestBody))
	w := httptest.NewRecorder()
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	var createUsecase = usecases.NewCreateUserUseCase(usersRepository)
	handler := api.CreateUserHandler{UseCase: createUsecase}
	res := w.Result()
	handler.ServeHTTP(w, req)
	defer res.Body.Close()
	var response api.JsonCreateUserResponse
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		log.Fatalln(err)
	}
	status := w.Code
	assert.Equal(t, http.StatusOK, status)
}

func TestCreatingUserShouldReturnStatus400(t *testing.T) {
	var requestBody = `{"username": "katya78", "birthday": "2006-12-13", "gender": 1}`
	req := httptest.NewRequest(http.MethodPost, "/api/user/create", strings.NewReader(requestBody))
	w := httptest.NewRecorder()
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	var createUsecase = usecases.NewCreateUserUseCase(usersRepository)
	handler := api.CreateUserHandler{UseCase: createUsecase}
	res := w.Result()
	handler.ServeHTTP(w, req)
	defer res.Body.Close()
	status := w.Code
	assert.Equal(t, http.StatusBadRequest, status)
}

func TestFindUserByIdShouldReturnStatus200(t *testing.T) {
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var testUser = user.NewUser("Steve", "Steve13", date, "+79528123333", user.Male)
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	_, err := usersRepository.Create(context.Background(), testUser)
	if err != nil {
		t.Fatalf("")
	}
	var responseBody = api.JsonFindUserByIdResponse{}
	responseBody.Username = "Steve"
	responseBody.DisplayName = "Steve13"
	responseBody.CurrentMeetingID = uuid.NullUUID{}
	responseBody.Rating = 0
	responseBody.Gender = user.Male
	responseBody.Birthday = date
	responseBody.ID = testUser.ID
	req := httptest.NewRequest(http.MethodGet, "/{userID}", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", testUser.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	var findUsecase = usecases.NewFindUserByIdUseCase(usersRepository)
	handler := api.FindUserByIdHandler{UseCase: findUsecase}
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	var response api.JsonFindUserByIdResponse
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		log.Fatalln(err)
	}
	status := w.Code
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, responseBody, response)
	usersRepository.Delete(context.Background(), testUser)
}

func TestFindUserByCriteriaShouldReturnStatus200(t *testing.T) {
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var testUser = user.NewUser("Steve", "Steve13", date, "+79528123333", user.Male)
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	_, err := usersRepository.Create(context.Background(), testUser)
	if err != nil {
		t.Fatalf("")
	}
	req := httptest.NewRequest(http.MethodGet, "/find?", nil)
	req.URL.Query().Add("user_name", testUser.Username)
	w := httptest.NewRecorder()
	req = mux.SetURLVars(req, map[string]string{"user_name": testUser.Username})
	var findUsecase = usecases.NewFindUsersByCriteriaUseCase(usersRepository)
	handler := api.FindUsersByCriteriaHandler{UseCase: findUsecase}
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	var response []user.User
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		log.Fatalln(err)
	}
	status := w.Code
	assert.Equal(t, http.StatusOK, status)
	var IDs = []uuid.UUID{}
	var userFromRepository = user.User{}
	for _, element := range response {
		IDs = append(IDs, element.ID)
		if element.ID == testUser.ID {
			userFromRepository = element
		}
	}
	assert.True(t, AreUsersEqual(testUser, &userFromRepository))
	usersRepository.Delete(context.Background(), testUser)
}

func TestUpdateUserShouldReturnStatus200(t *testing.T) {
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var testUser = user.NewUser("Katya", "Katya13", date, "+79528123333", user.Female)
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	_, err := usersRepository.Create(context.Background(), testUser)
	if err != nil {
		t.Fatalf("")
	}
	testUser.Username = "Katya78"
	usersRepository.Update(context.Background(), testUser)
	var requestBody = `{"username": "katya78", "gender": 1}`
	req := httptest.NewRequest(http.MethodPut, "/api/user/update/{user_id}", strings.NewReader(requestBody))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", testUser.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	var updateUsecase = usecases.NewUpdateUserUseCase(usersRepository)
	handler := api.UpdateUserHandler{UseCase: updateUsecase}
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	var responseID api.JsonUpdateUserResponse
	if err := json.Unmarshal([]byte(w.Body.String()), &responseID); err != nil {
		log.Fatalln(err)
	}
	status := w.Code
	assert.Equal(t, http.StatusOK, status)

	req = httptest.NewRequest(http.MethodGet, "/{userID}", nil)
	rctx = chi.NewRouteContext()
	rctx.URLParams.Add("userID", testUser.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w = httptest.NewRecorder()
	var findUsecase = usecases.NewFindUserByIdUseCase(usersRepository)
	newHandler := api.FindUserByIdHandler{UseCase: findUsecase}
	newHandler.ServeHTTP(w, req)
	res = w.Result()
	defer res.Body.Close()
	var response api.JsonFindUserByIdResponse
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		log.Fatalln(err)
	}
	var responseBody = api.JsonFindUserByIdResponse{}
	responseBody.Username = "katya78"
	/*responseBody.DisplayName = "Katya13"
	responseBody.CurrentMeetingID = uuid.NullUUID{}
	responseBody.Rating = 0
	responseBody.Birthday = date*/
	responseBody.Gender = user.Female
	responseBody.ID = testUser.ID
	assert.Equal(t, responseBody, response)
	usersRepository.Delete(context.Background(), testUser)
	// ок, если таки должно быть, что старые поля затираются нулями, если не поданы в запросе
}

func TestDeletingUser(t *testing.T) {
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var testUser = user.NewUser("Katya", "Katya13", date, "+79528123333", user.Female)
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	req := httptest.NewRequest(http.MethodDelete, "/api/user/update/{user_id}", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", testUser.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	var deleteUsecase = usecases.NewDeleteUserUseCase(usersRepository)
	handler := api.DeleteUserHandler{UseCase: deleteUsecase}
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	status := w.Code
	assert.Equal(t, http.StatusInternalServerError, status)

	_, err := usersRepository.Create(context.Background(), testUser)
	if err != nil {
		t.Fatalf("Error in creating user %v", err)
	}
	req = httptest.NewRequest(http.MethodDelete, "/api/user/update/{userID}", nil)
	rctx = chi.NewRouteContext()
	rctx.URLParams.Add("userID", testUser.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	res = w.Result()
	defer res.Body.Close()
	status = w.Code
	var responseID api.JsonUpdateUserResponse
	json.Unmarshal([]byte(w.Body.String()), &responseID)
	assert.Equal(t, http.StatusOK, status)
	assert.True(t, reflect.DeepEqual(responseID, api.JsonUpdateUserResponse{}))
	usersRepository.Delete(context.Background(), testUser)
}
