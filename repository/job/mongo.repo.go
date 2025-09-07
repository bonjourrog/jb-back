package job

import (
	"context"
	"errors"
	"os"

	"github.com/bonjourrog/jb/entity/job"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type JobRepository interface {
	Create(job job.Post, ctx context.Context) error
	GetAll(filter bson.M, page int, ctx context.Context) ([]job.PostWithCompany, int64, error)
	Update(job job.Post, ctx context.Context) error
	Delete(job_id bson.ObjectID, user_id bson.ObjectID, ctx context.Context) error
}

type jobRepository struct {
	client *mongo.Client
}

func NewJobRepository(client *mongo.Client) JobRepository {
	return &jobRepository{client: client}
}

func (r *jobRepository) Create(job job.Post, ctx context.Context) error {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("jobs")
	_, err := coll.InsertOne(ctx, job)
	if err != nil {
		return err
	}
	return nil
}
func (r *jobRepository) GetAll(filter bson.M, page int, ctx context.Context) ([]job.PostWithCompany, int64, error) {
	var (
		results     []job.PostWithCompany
		limit       = 10
		skip        = (page - 1) * limit
		lookupStage bson.D
	)
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("jobs")
	total, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	var userId = filter["user_id"]
	if userId != bson.NilObjectID {
		lookupStage = bson.D{{
			"$lookup", bson.M{
				"from": "applications",
				"let":  bson.M{"jobId": "$_id"},
				"pipeline": mongo.Pipeline{
					{{"$match", bson.M{
						"$expr": bson.M{
							"$and": bson.A{
								bson.M{"$eq": bson.A{"$job_id", "$$jobId"}},
								bson.M{"$eq": bson.A{"$user_id", userId}},
							},
						},
					}}},
				},
				"as": "userApplications",
			},
		}}
	} else {
		lookupStage = bson.D{}
	}
	delete(filter, "user_id")

	pipeline := mongo.Pipeline{
		lookupStage,
		{
			{"$match", bson.M{"userApplications": bson.M{"$size": 0}}},
		},
		// Filtrado
		{{"$match", filter}},

		// Paginación
		{{"$skip", int64(skip)}},
		{{"$limit", int64(limit)}},

		// Join con companies
		{{"$lookup", bson.M{
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
		}}},
	}

	// findOptions := options.Find()
	// findOptions.SetSkip(int64(skip)).SetLimit(int64(limit))
	// cursor, err := coll.Find(context.TODO(), filter, findOptions)
	cursor, err := coll.Aggregate(ctx, pipeline)
	// for cursor.Next(context.TODO()) {
	// 	var raw bson.M
	// 	cursor.Decode(&raw)
	// 	fmt.Printf("%+v\n", raw)
	// 	break
	// }
	if err != nil {
		return results, 0, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &results); err != nil {
		return results, 0, err
	}

	return results, total, nil
}

func (r *jobRepository) Update(job job.Post, ctx context.Context) error {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("jobs")
	_, err := coll.UpdateOne(ctx, bson.M{"_id": job.ID}, bson.M{"$set": job})
	if err != nil {
		return err
	}
	return nil
}
func (r *jobRepository) Delete(job_id bson.ObjectID, user_id bson.ObjectID, ctx context.Context) error {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("jobs")
	res, err := coll.DeleteOne(ctx, bson.M{"_id": job_id, "company_id": user_id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no se pudo eliminar el empleo")
	}
	return nil
}
