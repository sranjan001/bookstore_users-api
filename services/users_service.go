package services

import (
	"fmt"
	"github.com/sranjan001/bookstore_users-api/domain/users"
	"github.com/sranjan001/bookstore_users-api/util/date_utils"
	"github.com/sranjan001/bookstore_users-api/util/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestError) {
	fmt.Println("User id: ", userId)
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(userId int64) *errors.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}

func Search(status string) ([]users.User, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
