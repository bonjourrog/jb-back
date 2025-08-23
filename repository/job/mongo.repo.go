package job

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/bonjourrog/jb/entity/job"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type JobRepository interface {
	Create(job job.Post) error
	GetAll(filter bson.M, page int) ([]job.PostWithCompany, int64, error)
	Update(job job.Post) error
	Delete(job_id bson.ObjectID, user_id bson.ObjectID) error
	ApplyToJob(application job.Application) error
	IsUserAlreadyApplied(user_id bson.ObjectID, job_id bson.ObjectID) (bool, error)
}

type jobRepository struct {
	client *mongo.Client
}

func NewJobRepository(client *mongo.Client) JobRepository {
	return &jobRepository{client: client}
}

func (r *jobRepository) Create(job job.Post) error {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("jobs")
	_, err := coll.InsertOne(context.TODO(), job)
	if err != nil {
		return err
	}
	return nil
}
func (r *jobRepository) GetAll(filter bson.M, page int) ([]job.PostWithCompany, int64, error) {
	var (
		results     []job.PostWithCompany
		limit       = 10
		skip        = (page - 1) * limit
		lookupStage bson.D
	)
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("jobs")
	total, err := coll.CountDocuments(context.TODO(), filter)
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

	fmt.Println(userId)
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
	cursor, err := coll.Aggregate(context.TODO(), pipeline)
	// for cursor.Next(context.TODO()) {
	// 	var raw bson.M
	// 	cursor.Decode(&raw)
	// 	fmt.Printf("%+v\n", raw)
	// 	break
	// }
	if err != nil {
		return results, 0, err
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &results); err != nil {
		return results, 0, err
	}

	return results, total, nil
}

func (r *jobRepository) Update(job job.Post) error {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("jobs")
	_, err := coll.UpdateOne(context.TODO(), bson.M{"_id": job.ID}, bson.M{"$set": job})
	if err != nil {
		return err
	}
	return nil
}
func (r *jobRepository) Delete(job_id bson.ObjectID, user_id bson.ObjectID) error {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("jobs")
	res, err := coll.DeleteOne(context.TODO(), bson.M{"_id": job_id, "company_id": user_id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no se pudo eliminar el empleo")
	}
	return nil
}
func (r *jobRepository) ApplyToJob(application job.Application) error {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("applications")

	// Insert the application into the database
	_, err := coll.InsertOne(context.TODO(), application)
	if err != nil {
		return err
	}

	return nil
}

func (r *jobRepository) IsUserAlreadyApplied(user_id bson.ObjectID, job_id bson.ObjectID) (bool, error) {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("applications")

	count, err := coll.CountDocuments(context.TODO(), bson.M{
		"user_id": user_id,
		"job_id":  job_id,
	})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
