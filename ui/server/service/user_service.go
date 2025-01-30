package service

import (
	"errors"
	"html"
	"regexp"
	"strings"

	"forum/server/shareddata"

	"forum/server/data"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserData *data.UserData
}

func (s *UserService) RegisterUser(user *shareddata.User) error {
	// Username
	if len((*user).Username) < 3 || len((*user).Username) > 15 {
		return errors.New(shareddata.Errors.InvalidUsername)
	}
	// Password
	if len((*user).Password) < 6 || len((*user).Password) > 30 {
		return errors.New(shareddata.Errors.InvalidPassword)
	}
	// email
	(*user).Email = strings.ToLower((*user).Email)
	if EmailChecker((*user).Email) {
		return errors.New(shareddata.Errors.InvalidEmail)
	}
	if len((*user).Email) > 50 {
		return errors.New(shareddata.Errors.LongEmail)
	}
	// username or email existance
	if s.UserData.CheckIfUserExists((*user).Username, (*user).Email) {
		return errors.New(shareddata.Errors.UserAlreadyExist)
	}
	// Generate Uuid
	(*user).Uuid = GenerateUuid()
	// Encrypt Pass
	var err error
	(*user).Password, err = EncyptPassword((*user).Password)
	if err != nil {
		return err
	}

	// Fix username html
	(*user).Username = html.EscapeString((*user).Username)

	// Insert the user
	return s.UserData.InsertUser(*user)
}

func (s *UserService) LoginUser(user *shareddata.User) error {
	// email
	(*user).Email = strings.ToLower((*user).Email)
	if EmailChecker((*user).Email) {
		return errors.New(shareddata.Errors.InvalidEmail)
	}
	if len((*user).Email) > 50 {
		return errors.New(shareddata.Errors.LongEmail)
	}
	// Password
	if len((*user).Password) < 6 || len((*user).Password) > 30 {
		return errors.New(shareddata.Errors.InvalidPassword)
	}
	// check existance
	if !s.UserData.CheckIfUserExists((*user).Username, (*user).Email) {
		return errors.New(shareddata.Errors.InvalidCredentials)
	}
	// get user password
	UserPassword, err := s.UserData.GetUserPassword((*user).Email)
	if err != nil {
		return err
	}

	// Check Password Validity
	if !CheckPasswordValidity(UserPassword, (*user).Password) {
		return errors.New(shareddata.Errors.InvalidCredentials)
	}

	// generate new uuid
	(*user).Uuid = GenerateUuid()

	// Update uuid
	s.UserData.UpdateUuid((*user).Uuid, (*user).Email)

	return nil
}

// Helpers
func EmailChecker(email string) bool {
	return !regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString(email)
}

func GenerateUuid() string {
	return uuid.Must(uuid.NewV4()).String()
}

func EncyptPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func CheckPasswordValidity(hashedPass, entredPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(entredPass))
	return err == nil
}
