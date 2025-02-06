package usecase

import (
	"errors"

	"github.com/Ayyasy123/pt-aka-tech-vision/internal/model"
	"github.com/Ayyasy123/pt-aka-tech-vision/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	CreateUser(req *model.CreateUserReq) (*model.UserRes, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) CreateUser(req *model.CreateUserReq) (*model.UserRes, error) {
	// cek email sudah ada atau belum
	exist, err := u.userRepository.IsEmailExists(req.Email)
	if err != nil {
		return nil, err
	}

	if exist {
		return nil, errors.New("email already exists")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	err = u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &model.UserRes{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
