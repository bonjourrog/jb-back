package job

import (
	"context"
	"os"

	"github.com/bonjourrog/jb/db"
	"github.com/bonjourrog/jb/entity/job"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type JobRepository interface {
	Create(job job.Post) error
	GetAll(filter bson.M) ([]job.Post, error)
}

type jobRepository struct{}

func NewJobRepository() JobRepository {
	return &jobRepository{}
}

func (*jobRepository) Create(job job.Post) error {
	var (
		_db = db.NewMongoConnection()
	)

	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	coll := client.Database(os.Getenv("DATABASE")).Collection("jobs")
	_, err := coll.InsertOne(context.TODO(), job)
	if err != nil {
		return err
	}
	return nil
}
func (*jobRepository) GetAll(filter bson.M) ([]job.Post, error) {
	var (
		_db     = db.NewMongoConnection()
		results []job.Post
	)

	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	coll := client.Database(os.Getenv("DATABASE")).Collection("jobs")
	findOptions := options.Find()
	findOptions.SetLimit(12)
	cursor, err := coll.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return results, err
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &results); err != nil {
		return results, err
	}

	return results, nil
}
