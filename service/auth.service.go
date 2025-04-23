package service

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/bonjourrog/jb/entity"
	"github.com/bonjourrog/jb/repository/auth"
	"github.com/bonjourrog/jb/util"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthService interface {
	Signup(user entity.User) (*mongo.InsertOneResult, error)
	SignIn(credentials entity.Account) (string, error)
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

	userFound, err := _authRepository.FindByEmail(user.Account.Email)
	if err != nil {
		return nil, err
	}
	if userFound != nil {
		return nil, errors.New("email already exists")
	}
	return _authRepository.Create(user)
}
func (*authService) SignIn(credentials entity.Account) (string, error) {
	user, err := _authRepository.FindByEmail(credentials.Email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("incorrect email")
	}
	if strings.TrimSpace(credentials.Password) == "" || strings.TrimSpace(credentials.Email) == "" {
		return "", errors.New("some required fields are empty")
	}
	if ok := util.ComparePassword([]byte(user.Account.Password), []byte(credentials.Password)); !ok {
		return "", errors.New("password incorrect")
	}
	claims := entity.CustomeClaims{
		Role:   user.Role,
		Email:  user.Account.Email,
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mySigningKey := os.Getenv("SigningKey")
	ss, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		return "", err
	}

	return ss, nil
}
