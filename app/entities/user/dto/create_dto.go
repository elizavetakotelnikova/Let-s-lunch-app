package dto

import "cmd/app/entities/user"

type CreateUserDto struct {
	Username    string      `json:"username"`
	DisplayName string      `json:"displayName"`
	Age         int         `json:"age"`
	Gender      user.Gender `json:"gender"`
}
