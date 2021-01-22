package user

import (
	userModel "github.com/smart-evolution/shapi/domain/models/user"
)

// IStateRepository - user repository interface
type IStateRepository interface {
	SetIsAlerts(i bool)
	IsAlerts() bool
	SetSendAlert(i bool)
	SendAlert() bool
}

// IUserRepository - user repository interface
type IUserRepository interface {
	FindOneUser(query interface{}) (userModel.User, error)
	FindAllUsers(query interface{}) ([]userModel.User, error)
	RegisterUser(user userModel.User) error
}
