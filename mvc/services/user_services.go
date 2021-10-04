package services

import (
	"github.com/mayuka-c/golang-microservices/mvc/domain"
	"github.com/mayuka-c/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
