package service

import (
	"bookapi/entity"
	"errors"
)

var Users []entity.User

func InitializeUsers() {
	Users = []entity.User{}
}
func GetAllUsers() []entity.User {
	return Users
}

func Register(user entity.User) entity.User {
	user.ID = uint64(len(Users) + 1)
	Users = append(Users, user)
	return user
}

func Profile(userID uint64) (entity.User, error) {
	for i := 0; i < len(Users); i++ {
		if Users[i].ID == userID {
			return Users[i], nil
		}
	}
	return entity.User{}, errors.New("The specified user does not exist!")
}

func UpdateProfile(user entity.User) (entity.User, error) {
	for i := 0; i < len(Users); i++ {
		if Users[i].ID == user.ID {
			Users[i] = user
			return user, nil
		}
	}
	return user, errors.New("The specified user does not exist!")
}

func DeleteAccount(userID uint64) error {
	for i := 0; i < len(Users); i++ {
		if Users[i].ID == userID {
			Users = append(Users[:i], Users[i+1:]...)
			return nil
		}
	}
	return errors.New("The specified user does not exist!")
}
