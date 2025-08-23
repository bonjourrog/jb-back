package service

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/bonjourrog/jb/entity"
	"github.com/bonjourrog/jb/repository/auth"
	"github.com/bonjourrog/jb/util"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthService interface {
	Signup(user entity.User, ctx context.Context) (*mongo.InsertOneResult, error)
	SignIn(credentials entity.Account, ctx context.Context) (string, error)
}

type authService struct{}

var (
	_authRepository auth.AuthRepo
)

func NewAuthService(authRepository auth.AuthRepo) AuthService {
	_authRepository = authRepository
	return &authService{}
}

func (*authService) Signup(user entity.User, ctx context.Context) (*mongo.InsertOneResult, error) {
	//Remove leading and trailing spaces from some fields in case they have them
	user.Account.Email = strings.TrimSpace(strings.ToLower(user.Account.Email))
	user.Name = strings.TrimSpace(strings.ToLower(user.Name))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.Company.Address.FirstStreet = strings.TrimSpace(strings.ToLower(user.Company.Address.FirstStreet))
	user.Company.Address.SecondStreet = strings.TrimSpace(strings.ToLower(user.Company.Address.SecondStreet))
	user.Company.Address.Neighborhood = strings.TrimSpace(strings.ToLower(user.Company.Address.Neighborhood))

	userFound, err := _authRepository.FindByEmail(user.Account.Email, ctx)

	if err != nil {
		return nil, err
	}
	if userFound != nil {
		return nil, errors.New("email already exists")
	}
	//Validate if signup is for a company role
	if user.Role == entity.RoleCompany {
		if user.Company.Name == "" || user.Company.Logo == "" || len(user.Company.Address.Location.Coordinates) != 2 || user.Company.Address.Location.Type != "Point" {
			return nil, errors.New("some required fields are empty")
		}
	}
	if user.Name == "" || user.LastName == "" || user.Account.Email == "" || user.Account.Password == "" {
		return nil, errors.New("some required fields are empty")
	}
	if isRoleValid := util.VerifyRole(user.Role); !isRoleValid {
		return nil, errors.New("invalid role")
	}
	hashedPassword, err := util.GeneratePassword(user.Account.Password)
	if err != nil {
		return nil, err
	}

	user.ID = bson.NewObjectID()
	user.Account.Password = hashedPassword
	user.Account.Banned = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return _authRepository.Create(user, ctx)
}
func (*authService) SignIn(credentials entity.Account, ctx context.Context) (string, error) {
	user, err := _authRepository.FindByEmail(credentials.Email, ctx)
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
