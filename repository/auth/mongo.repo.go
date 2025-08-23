package auth

import (
	"context"
	"os"

	"github.com/bonjourrog/jb/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type authrepo struct {
	client *mongo.Client
}

type AuthRepo interface {
	Create(user entity.User, ctx context.Context) (*mongo.InsertOneResult, error)
	FindByEmail(email string, ctx context.Context) (*entity.User, error)
}

func NewAuthRepository(client *mongo.Client) AuthRepo {
	return &authrepo{client: client}
}

// Create insert a new user in dadabase
func (r *authrepo) Create(user entity.User, ctx context.Context) (*mongo.InsertOneResult, error) {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("users")
	insertResult, err := coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}

// FindByEmail checks if the email exists in the database
func (r *authrepo) FindByEmail(email string, ctx context.Context) (*entity.User, error) {
	var (
		user entity.User
	)
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("users")
	err := coll.FindOne(ctx, bson.M{"account.email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
