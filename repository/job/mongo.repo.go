package job

import (
	"context"
	"os"

	"github.com/bonjourrog/jb/db"
	"github.com/bonjourrog/jb/entity/job"
)

type JobRepository interface {
	Create(job job.Post) error
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
