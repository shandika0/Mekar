package services

import (
	"errors"
	"fmt"
	"pretty/models"
	"pretty/repo"
)

type UserServiceStruct struct {
	userRepo repo.UserRepoInterface
}

type UserServiceInterface interface {
	AddUser(user *models.User) (*models.User, error)
	GetAllUser() (*[]models.User, error)
	GetUserByID(id int) (*models.User, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
}

func CreateUserServiceImpl(userRepo repo.UserRepoInterface) UserServiceInterface {
	return &UserServiceStruct{userRepo}
}

func (u *UserServiceStruct) AddUser(user *models.User) (*models.User, error) {
	data, err := u.userRepo.AddUser(user)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserServiceStruct) GetAllUser() (*[]models.User, error) {
	data, err := u.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return data, nil
}

func (u *UserServiceStruct) GetUserByID(id int) (*models.User, error) {
	result, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserServiceStruct) UpdateUser(id int, user *models.User) (*models.User, error) {
	_, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("id does not exist")
	}
	result, err := u.userRepo.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserServiceStruct) DeleteUser(id int) error {

	_, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return errors.New("ID Admin does not exist")
	}
	err = u.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
