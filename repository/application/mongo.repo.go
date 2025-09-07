package application

import (
	"context"
	"os"

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

func (r *applicationRepo) GetByIds(job_ids []bson.ObjectID, ctx context.Context) ([]job.PostWithCompany, error) {
	var (
		results []job.PostWithCompany
		err     error
	)
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("jobs")
	cursor, err := coll.Find(ctx, bson.M{"_id": bson.M{"$in": job_ids}})
	// pipeline := mongo.Pipeline{
	// {{"$match", bson.D{{"_id", bson.D{{"$in", job_ids}}}}}},

	// {{
	// 	"$lookup", bson.M{
	// 		"from":         "users",
	// 		"localField":   "company_id",
	// 		"foreignField": "_id",
	// 		"as":           "company",
	// 	}}},
	// {{"$unwind", bson.M{
	// 	"path":                       "$company",
	// 	"preserveNullAndEmptyArrays": true,
	// }}},
	// {"$addFields", bson.M{
	// 	"company_name": "$company.company.name",
	// 	"company_logo": "$company.company.logo",
	// }},

	// }

	// cursor, err := coll.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}
