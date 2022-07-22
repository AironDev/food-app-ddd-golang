package application

import (
	"github.com/airondev/food-app-ddd-golang/domain/entity"
	"github.com/airondev/food-app-ddd-golang/domain/repository"
)

type userApp struct {
	model repository.UserRepository
}

//UserApp implements the UserAppInterface
var _ = &userApp{}

type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUsers() ([]entity.User, error)
	GetUser(uint64) (*entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.model.SaveUser(user)
}

func (u *userApp) GetUsers() ([]entity.User, error) {
	return u.model.GetUsers()
}

func (u *userApp) GetUser(userId uint64) (*entity.User, error) {
	return u.model.GetUser(userId)
}

func (u *userApp) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return u.model.GetUserByEmailAndPassword(user)
}
