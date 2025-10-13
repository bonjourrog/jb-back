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
	Delete(job_id bson.ObjectID, company_id bson.ObjectID, ctx context.Context) error
	GetById(job_id bson.ObjectID, ctx context.Context) (*job.Post, error)
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
		limit       = 12
		skip        = (page - 1) * limit
		lookupStage bson.D
	)
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("jobs")
	// total, err := coll.CountDocuments(ctx, filter)
	// if err != nil {
	// 	return nil, 0, err
	// }
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

	basePipeline := mongo.Pipeline{}
	if len(lookupStage) > 0 {
		basePipeline = append(basePipeline, lookupStage)
		basePipeline = append(basePipeline, bson.D{{"$match", bson.M{"userApplications": bson.M{"$size": 0}}}})
	}

	if len(filter) > 0 {
		basePipeline = append(basePipeline, bson.D{{"$match", filter}})
	}

	countPipeline := append(basePipeline, bson.D{{"$count", "total"}})

	countCursor, err := coll.Aggregate(ctx, countPipeline)
	if err != nil {
		return nil, 0, err
	}
	var countResult []bson.M
	if err = countCursor.All(ctx, &countResult); err != nil {
		return nil, 0, err
	}

	var totalCount int64 = 0
	if len(countResult) > 0 {
		totalCount = int64(countResult[0]["total"].(int32))
	}

	// pipeline := mongo.Pipeline{
	// 	lookupStage,
	// 	{
	// 		{"$match", bson.M{"userApplications": bson.M{"$size": 0}}},
	// 	},
	// 	// Filtrado
	// 	{{"$match", filter}},

	// 	// Paginación
	// 	{{"$skip", int64(skip)}},
	// 	{{"$limit", int64(limit)}},

	// 	// Join con companies
	// 	{{"$lookup", bson.M{
	// 		"from":         "users",
	// 		"localField":   "company_id",
	// 		"foreignField": "_id",
	// 		"as":           "company",
	// 	}}},
	// 	{{"$unwind", bson.M{
	// 		"path":                       "$company",
	// 		"preserveNullAndEmptyArrays": true,
	// 	}}},
	// 	{{"$addFields", bson.M{
	// 		"company_name": "$company.company.name",
	// 		"company_logo": "$company.company.logo",
	// 	}}},
	// }

	// cursor, err := coll.Aggregate(ctx, basePipeline)

	// if err != nil {
	// 	return results, 0, err
	// }
	// defer cursor.Close(ctx)
	// if err = cursor.All(ctx, &results); err != nil {
	// 	return results, 0, err
	// }

	dataPipeline := append(basePipeline,
		bson.D{{"$skip", int64(skip)}},
		bson.D{{"$limit", int64(limit)}},

		// Join con companies
		bson.D{{"$lookup", bson.M{
			"from":         "users",
			"localField":   "company_id",
			"foreignField": "_id",
			"as":           "company",
		}}},
		bson.D{{"$unwind", bson.M{
			"path":                       "$company",
			"preserveNullAndEmptyArrays": true,
		}}},
		bson.D{{"$addFields", bson.M{
			"company_name":  "$company.company.name",
			"company_logo":  "$company.company.logo",
			"company_phone": "$company.account.phone",
		}}},
	)

	cursor, err := coll.Aggregate(ctx, dataPipeline)
	if err != nil {
		return results, 0, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &results); err != nil {
		return results, 0, err
	}

	return results, totalCount, nil
}

func (r *jobRepository) Update(job job.Post, ctx context.Context) error {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("jobs")
	_, err := coll.UpdateOne(ctx, bson.M{"_id": job.ID}, bson.M{"$set": job})
	if err != nil {
		return err
	}
	return nil
}
func (r *jobRepository) Delete(job_id bson.ObjectID, company_id bson.ObjectID, ctx context.Context) error {
	coll := r.client.Database(os.Getenv("DATABASE")).Collection("jobs")
	res, err := coll.DeleteOne(ctx, bson.M{"_id": job_id, "company_id": company_id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no se pudo eliminar el empleo")
	}
	return nil
}
func (r *jobRepository) GetById(job_id bson.ObjectID, ctx context.Context) (*job.Post, error) {
	coll := r.client.Database((os.Getenv("DATABASE"))).Collection("jobs")
	var result job.Post
	err := coll.FindOne(ctx, bson.M{"_id": job_id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
