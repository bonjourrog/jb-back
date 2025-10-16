package user

import (
	"context"
	"os"

	"github.com/bonjourrog/jb/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepo interface {
	FindByField(field string, value interface{}, ctx context.Context) (*entity.User, error)
}
type userRepo struct {
	client *mongo.Client
}

func NewUserRepo(client *mongo.Client) UserRepo {
	return &userRepo{
		client: client,
	}
}
func (u *userRepo) FindByField(field string, value interface{}, ctx context.Context) (*entity.User, error) {
	var (
		user entity.User
	)
	coll := u.client.Database((os.Getenv("DATABASE"))).Collection("users")
	if err := coll.FindOne(ctx, bson.M{field: value}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
