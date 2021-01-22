package user

import (
	"errors"
	goutilsSession "github.com/coda-it/goutils/session"
	"github.com/coda-it/gowebserver/session"
	"github.com/coda-it/gowebserver/utils/logger"
	"github.com/smart-evolution/shapi/constants"
	userModel "github.com/smart-evolution/shapi/domain/models/user"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

// Usecase - user usecases
type Usecase struct {
	stateRepository IStateRepository
	userRepository  IUserRepository
}

// New - creates new user usecases
func New(sr IStateRepository, ur IUserRepository) *Usecase {
	return &Usecase{
		sr,
		ur,
	}
}

// AuthenticateByCredentials - authenticate user with credentials
func (u *Usecase) AuthenticateByCredentials(username string, password string) (userModel.User, error) {
	usr, err := u.userRepository.FindOneUser(bson.M{
		"username": username,
		"password": password,
	})

	if err != nil {
		msg := "user not found"
		logger.Log(msg, logger.ERROR)
		return userModel.User{}, errors.New(msg)
	}

	logger.Log("logged in as user", usr.Username)

	return usr, nil
}

// CreateClientSession - authenticate uer with credentials and create session cookie
func (u *Usecase) CreateClientSession(w http.ResponseWriter, r *http.Request, username string, password string, sm session.ISessionManager) bool {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	authenticatedUser, err := u.AuthenticateByCredentials(username, password)

	if err == nil {
		t := time.Now()
		timeStr := t.Format(time.RFC850)
		cookieValue := goutilsSession.CreateSessionID(username, password, timeStr)

		cookie := http.Cookie{
			Name:    constants.SessionKey,
			Value:   cookieValue,
			Expires: expiration}

		session := sm.Create(cookieValue)
		session.Set("user", authenticatedUser)

		http.SetCookie(w, &cookie)
		return true
	}
	return false
}

// RegisterUser - registers user
func (u *Usecase) RegisterUser(user userModel.User) error {
	return u.userRepository.RegisterUser(user)
}

// FindAllUsers - gets all registered users
func (u *Usecase) FindAllUsers() ([]userModel.User, error) {
	return u.userRepository.FindAllUsers(bson.M{})
}

// SetIsAlerts - sets alerts
func (u *Usecase) SetIsAlerts(i bool) {
	u.stateRepository.SetIsAlerts(i)
}

// IsAlerts - gets are alerts turned on
func (u *Usecase) IsAlerts() bool {
	return u.stateRepository.IsAlerts()
}

// SetSendAlert - sets should alerts be sent
func (u *Usecase) SetSendAlert(i bool) {
	u.stateRepository.SetSendAlert(i)
}

// SendAlert - sends alerts
func (u *Usecase) SendAlert() bool {
	return u.stateRepository.SendAlert()
}
