package repository

import (
	"github.com/Ayyasy123/pt-aka-tech-vision/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	IsEmailExists(email string) (bool, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(req *model.User) error {
	return r.db.Create(&req).Error
}

func (r *userRepo) IsEmailExists(email string) (bool, error) {
	var user model.User
	r.db.Where("email = ?", email).First(&user)
	return user.ID > 0, nil
}
