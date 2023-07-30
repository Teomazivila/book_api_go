package service

import (
	"bookapi/dto"
	"bookapi/entity"
	"bookapi/repository"
	"errors"
	"log"

	"github.com/mashingan/smapping"
)

func GetAllUsers() []entity.User {
	return repository.GetAllUsers()
}

func InsertUsers(userDTO dto.RegisterDTO, userID uint64) dto.UserResponseDTO {
	user := entity.User{}
	userResponse := dto.UserResponseDTO{}

	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		log.Fatal("failed to map ", err)
		return userResponse
	}

	user.ID = userID
	user = repository.InsertUser(user)

	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		log.Fatal("failed to map to response ", err)
		return userResponse
	}

	return userResponse
}

func Profile(userID uint64) (entity.User, error) {
	if user, err := repository.GetUser(userID); err == nil {
		return user, nil
	}
	return entity.User{}, errors.New("User does not exist!")
}

func UpdateProfile(userDTO dto.UserUpdateDTO, userID uint64) dto.UserResponseDTO {
	user := entity.User{}
	userResponse := dto.UserResponseDTO{}

	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		log.Fatal("failed to map ", err)
		return userResponse
	}

	user.ID = userID
	user = repository.UpdateUser(user)

	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		log.Fatal("failed to map to response ", err)
		return userResponse
	}

	return userResponse
}

func UpdateProfilse(user entity.User) (entity.User, error) {
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
