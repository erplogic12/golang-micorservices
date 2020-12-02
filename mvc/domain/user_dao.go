package domain

import (
	"fmt"
	"github.com/erplogic12/golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Gil", LastName: "Johnson", Email: "betos.depaul@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {

	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found bitch. Go check another database", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
