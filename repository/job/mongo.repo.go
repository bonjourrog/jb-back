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
	GetAll(filter bson.M, page int) ([]job.Post, int64, error)
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
func (*jobRepository) GetAll(filter bson.M, page int) ([]job.Post, int64, error) {
	var (
		_db     = db.NewMongoConnection()
		results []job.Post
		limit   = 10
		skip    = (page - 1) * limit
	)

	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	coll := client.Database(os.Getenv("DATABASE")).Collection("jobs")
	total, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, 0, err
	}
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip)).SetLimit(int64(limit))
	cursor, err := coll.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return results, 0, err
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &results); err != nil {
		return results, 0, err
	}

	return results, total, nil
}
