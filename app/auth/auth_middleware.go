package auth

import (
	users_repository "cmd/app/entities/user/repository"
	"context"
	"net/http"
)

type Config struct {
	Users users_repository.UsersRepository
}

func (a *Config) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), "user", "123")
		//fmt.Println(r.Header.Get("Authorization"))
		//
		//tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
		//
		////_, err := tokenAuth.Decode(r.Header.Get("Authorization"))
		////if err != nil {
		////	panic("aaaaaaa")
		////}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
