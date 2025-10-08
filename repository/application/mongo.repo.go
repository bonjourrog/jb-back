package application

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/bonjourrog/jb/entity"
	"github.com/bonjourrog/jb/entity/application"
	"github.com/bonjourrog/jb/entity/job"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ApplicationRepository interface {
	// Create(application application, ctx context.Context) error
	FindByUser(user_id bson.ObjectID, ctx context.Context) ([]application.Application, error)
	UpdateStatus(application_id bson.ObjectID, status string, ctx context.Context) error
	GetByIds(job_ids []bson.ObjectID, ctx context.Context) ([]job.PostWithCompany, error)
	FindByField(field string, value interface{}, ctx context.Context) (*[]application.Application, error)
	GetById(applicationId bson.ObjectID, ctx context.Context) (*application.Application, error)
	ApplyToJob(application application.Application, ctx context.Context) error
	IsUserAlreadyApplied(user_id bson.ObjectID, job_id bson.ObjectID, ctx context.Context) (bool, error)
	DeleteById(application_id bson.ObjectID, userId bson.ObjectID, ctx context.Context) error
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

func (a *applicationRepo) GetByIds(job_ids []bson.ObjectID, ctx context.Context) ([]job.PostWithCompany, error) {
	var (
		results []job.PostWithCompany
		err     error
	)
	coll := a.client.Database(os.Getenv("DATABASE")).Collection("jobs")
	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"_id", bson.D{{"$in", job_ids}}}}}},

		{{
			"$lookup", bson.M{
				"from":         "users",
				"localField":   "company_id",
				"foreignField": "_id",
				"as":           "company",
			}}},
		{{"$unwind", bson.M{
			"path":                       "$company",
			"preserveNullAndEmptyArrays": true,
		}}},
		{{"$addFields", bson.M{
			"company_name": "$company.company.name",
			"company_logo": "$company.company.logo",
		}},
		}}

	cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}
func (a *applicationRepo) GetById(applicationId bson.ObjectID, ctx context.Context) (*application.Application, error) {
	coll := a.client.Database((os.Getenv("DATABASE"))).Collection("applications")

	var app application.Application
	err := coll.FindOne(ctx, bson.M{"_id": applicationId}).Decode(&app)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, entity.ErrApplicationNotFound
		}
		return nil, err
	}
	return &app, nil
}
func (a *applicationRepo) ApplyToJob(application application.Application, ctx context.Context) error {
	coll := a.client.Database(os.Getenv("DATABASE")).Collection("applications")

	// Insert the application into the database
	_, err := coll.InsertOne(ctx, application)
	if err != nil {
		return err
	}

	return nil
}
func (a *applicationRepo) IsUserAlreadyApplied(user_id bson.ObjectID, job_id bson.ObjectID, ctx context.Context) (bool, error) {
	coll := a.client.Database(os.Getenv("DATABASE")).Collection("applications")

	count, err := coll.CountDocuments(ctx, bson.M{
		"user_id": user_id,
		"job_id":  job_id,
	})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
func (a *applicationRepo) DeleteById(application_id bson.ObjectID, userId bson.ObjectID, ctx context.Context) error {
	coll := a.client.Database(os.Getenv("DATABASE")).Collection("applications")

	result, err := coll.DeleteOne(ctx, bson.M{"_id": application_id, "user_id": userId})
	if err != nil {
		return fmt.Errorf("failed to delete application: %w", err)
	}
	if result.DeletedCount == 0 {
		return entity.ErrApplicationNotFound
	}

	return nil
}
func (a *applicationRepo) FindByField(field string, value interface{}, ctx context.Context) (*[]application.Application, error) {
	var results []application.Application
	coll := a.client.Database(os.Getenv("DATABASE")).Collection("applications")

	cursor, err := coll.Find(ctx, bson.M{field: value})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return &results, nil
}
