package job

import (
	"context"
	"os"

	"github.com/bonjourrog/jb/db"
	"github.com/bonjourrog/jb/entity/job"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type JobRepository interface {
	Create(job job.Post) error
	GetAll(filter bson.M, page int) ([]job.PostWithCompany, int64, error)
	Update(job job.Post) error
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
func (*jobRepository) GetAll(filter bson.M, page int) ([]job.PostWithCompany, int64, error) {
	var (
		_db     = db.NewMongoConnection()
		results []job.PostWithCompany
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
	pipeline := mongo.Pipeline{
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

func (*jobRepository) Update(job job.Post) error {
	var (
		_db = db.NewMongoConnection()
	)

	client := _db.Connection()
	defer func() {
		client.Disconnect(context.TODO())
	}()
	coll := client.Database(os.Getenv("DATABASE")).Collection("jobs")
	_, err := coll.UpdateOne(context.TODO(), bson.M{"_id": job.ID}, bson.M{"$set": job})
	if err != nil {
		return err
	}
	return nil
}
