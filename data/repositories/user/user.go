package user

import (
	"errors"
	"github.com/coda-it/gowebserver/utils/logger"
	"github.com/smart-evolution/shapi/data/persistence"
	userModel "github.com/smart-evolution/shapi/domain/models/user"
)

const (
	collectionName = "users"
)

// Repository - user repository
type Repository struct {
	Persistence persistence.IPersistance
}

// New - creates instance of user repository
func New(p persistence.IPersistance) *Repository {
	return &Repository{
		p,
	}
}

// FindOneUser - find one user entry
func (r *Repository) FindOneUser(query interface{}) (userModel.User, error) {
	c := r.Persistence.GetCollection(collectionName)

	var u userModel.User
	err := c.Find(query).One(&u)

	if err != nil {
		msg := "object not found"
		logger.Log(msg, logger.ERROR)
		return userModel.User{}, errors.New(msg)
	}

	return u, nil
}

// FindAllUsers - find all users
func (r *Repository) FindAllUsers(query interface{}) ([]userModel.User, error) {
	c := r.Persistence.GetCollection(collectionName)

	var users []userModel.User
	err := c.Find(query).All(&users)

	if err != nil {
		msg := "objects not found"
		logger.Log(msg, logger.ERROR)
		return nil, errors.New(msg)
	}

	return users, nil
}

// RegisterUser -
func (r *Repository) RegisterUser(user userModel.User) error {
	c := r.Persistence.GetCollection(collectionName)
	return c.Insert(user)
}
