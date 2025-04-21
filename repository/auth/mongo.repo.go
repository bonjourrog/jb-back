package auth

import (
	"context"
	"os"

	"github.com/bonjourrog/jb/db"
	"github.com/bonjourrog/jb/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type authrepo struct{}

type AuthRepo interface {
	Create(user entity.User) (*mongo.InsertOneResult, error)
	FindByEmail(email string) (*entity.User, error)
}

func NewAuthRepository() AuthRepo {
	return &authrepo{}
}

// Create insert a new user in dadabase
func (*authrepo) Create(user entity.User) (*mongo.InsertOneResult, error) {
	var (
		_db = db.NewMongoConnection()
	)
	client := _db.Connection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database(os.Getenv("DATABASE")).Collection("users")
	insertResult, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}

// FindByEmail checks if the email exists in the database
func (*authrepo) FindByEmail(email string) (*entity.User, error) {
	var (
		_db  = db.NewMongoConnection()
		user entity.User
	)
	client := _db.Connection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(os.Getenv("DATABASE")).Collection("users")
	err := coll.FindOne(context.TODO(), bson.M{"account.email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
