package repository

import (
	"bookapi/config"
	"bookapi/entity"
	"errors"
)

func InsertUser(user entity.User) entity.User {
	config.Db.Save(&user)
	config.Db.Preload("User").Find(&user)

	return user
}

func GetAllUsers() []entity.User {
	var users []entity.User
	config.Db.Preload("User").Find(&users)

	return users
}

func GetUser(userID uint64) (entity.User, error) {
	var user entity.User
	config.Db.Preload("User").First(&user, userID)
	if user.ID != 0 {
		return user, nil
	}

	return user, errors.New("requested user does not exists")
}

func UpdateUser(user entity.User) (entity.User, error) {
	if _, err := GetUser(user.ID); err == nil {
		config.Db.Save(&user)
		config.Db.Preload("User").Find(&user)
		return user, nil
	}
	return user, errors.New("requested user does not exists")
}

func DeleteUser(userID uint64) error {
	var user entity.User
	config.Db.First(&user, userID)
	if user.ID != 0 {
		config.Db.Delete(&user)
		return nil
	}
	return errors.New("requested user does not exists")
}

func GetTheUserkUsingID(userID uint64) entity.User {
	var user entity.User
	config.Db.Preload("User").First(&user, userID)
	return user
}
