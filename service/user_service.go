package service

import (
	"bookapi/entity"
	"bookapi/repository"
	"errors"
)

func GetAllUsers() []entity.User {
	return repository.GetAllUsers()
}

func InsertUsers(user entity.User) entity.User {
	user.ID = 2
	user = repository.InsertUser(user)
	return user
}

func Profile(userID uint64) (entity.User, error) {
	if user, err := repository.GetUser(userID); err == nil {
		return user, nil
	}
	return entity.User{}, errors.New("User does not exist!")
}

func UpdateProfile(user entity.User) (entity.User, error) {
	user.ID = 2
	if user, err := repository.UpdateUser(user); err == nil {
		return user, nil
	}
	return user, errors.New("User does not exist!")
}

func DeleteAccount(userID uint64) error {
	if err := repository.DeleteUser(userID); err == nil {
		return nil
	}
	return errors.New("User does not exist!")
}
