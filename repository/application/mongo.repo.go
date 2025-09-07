package application

import (
	"context"
	"os"

	"github.com/bonjourrog/jb/entity/application"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ApplicationRepository interface {
	// Create(application application, ctx context.Context) error
	FindByUser(user_id bson.ObjectID, ctx context.Context) ([]application.Application, error)
	UpdateStatus(application_id bson.ObjectID, status string, ctx context.Context) error
}

type applicationRepo struct {
	client *mongo.Client
}

func NewApplicationRepository(client *mongo.Client) ApplicationRepository {
	return &applicationRepo{
		client: client,
	}
}

func (a *applicationRepo) FindByUser(user_id bson.ObjectID, ctx context.Context) ([]application.Application, error) {
	var (
		results []application.Application
		err     error
		cursor  *mongo.Cursor
	)
	coll := a.client.Database(os.Getenv("DATABASE")).Collection("applications")
	cursor, err = coll.Find(ctx, bson.M{"user_id": user_id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (*applicationRepo) UpdateStatus(application_id bson.ObjectID, status string, ctx context.Context) error {
	return nil
}
