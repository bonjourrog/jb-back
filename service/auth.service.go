package service

import (
	"github.com/bonjourrog/jb/entity"
	"github.com/bonjourrog/jb/repository/auth"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthService interface {
	Signup(user entity.User) (*mongo.InsertOneResult, error)
}

type authService struct{}

var (
	_authRepository auth.AuthRepo
)

func NewAuthService(authRepository auth.AuthRepo) AuthService {
	_authRepository = authRepository
	return &authService{}
}

func (*authService) Signup(user entity.User) (*mongo.InsertOneResult, error) {
	if err := _authRepository.FindByEmail(user.Account.Email); err != nil {
		return nil, err
	}
	return _authRepository.Create(user)
}
