package domain

import (
	"fmt"
	"net/http"

	"github.com/mayuka-c/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "Mayuka", LastName: "C", Email: "mayuka1999@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}
