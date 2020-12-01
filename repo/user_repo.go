package repo

import (
	"fmt"
	"pretty/models"

	"github.com/jinzhu/gorm"
)

type UserRepoStruct struct {
	db *gorm.DB
}

type UserRepoInterface interface {
	AddUser(user *models.User) (*models.User, error)
	GetAllUser() (*[]models.User, error)
	GetUserByID(id int) (*models.User, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
}

func CreateUserRepoImpl(db *gorm.DB) UserRepoInterface {
	return &UserRepoStruct{db}
}

func (u *UserRepoStruct) AddUser(user *models.User) (*models.User, error) {
	tx := u.db.Begin()

	err := tx.Debug().Create(user).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("ERROR Add User")
		return nil, err
	}

	tx.Commit()
	return user, nil
}

func (u *UserRepoStruct) GetAllUser() (*[]models.User, error) {
	var user []models.User
	tx := u.db.Begin()

	err := tx.Debug().Preload("Pekerjaan").Preload("Pendidikan").Find(&user).Error
	if err != nil {
		fmt.Println("Error GetAllUser")
		return nil, err
	}

	return &user, nil
}

func (u *UserRepoStruct) GetUserByID(id int) (*models.User, error) {
	tx := u.db.Begin()
	data := models.User{}

	err := tx.Debug().Where("id = ?", id).Find(&data).Error
	if err != nil {
		fmt.Println("Error GetById")
		return nil, err
	}

	return &data, nil
}

func (u *UserRepoStruct) UpdateUser(id int, user *models.User) (*models.User, error) {
	tx := u.db.Begin()
	err := u.db.Debug().Model(&user).Where("id = ?", id, 1).Update(user).Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[UserRepo.Update] Error when query update data with error: %w", err)
	}

	tx.Commit()
	return user, nil
}

func (u *UserRepoStruct) DeleteUser(id int) error {
	data := models.User{}
	err := u.db.Debug().Where("id in (?)", id).Unscoped().Delete(&data).Error
	if err != nil {
		return fmt.Errorf("[UserRepo.Delete] Error when query delete data with error: %w", err)
	}

	return nil
}
