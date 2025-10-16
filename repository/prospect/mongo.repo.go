package prospect

import (
	"context"
	"os"

	"github.com/bonjourrog/jb/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ProspectRepo interface {
	InsertOne(prospect entity.Prospect, ctx context.Context) error
	FindByField(field string, value interface{}, ctx context.Context) (*entity.Prospect, error)
}

type prospectRepo struct {
	client *mongo.Client
}

func NewProspectRepo(client *mongo.Client) ProspectRepo {
	return &prospectRepo{client: client}
}
func (p *prospectRepo) InsertOne(prospect entity.Prospect, ctx context.Context) error {
	coll := p.client.Database(os.Getenv("DATABASE")).Collection("prospects")
	_, err := coll.InsertOne(ctx, prospect)
	if err != nil {
		return err
	}

	return nil
}
func (p *prospectRepo) FindByField(field string, value interface{}, ctx context.Context) (*entity.Prospect, error) {
	var (
		result entity.Prospect
	)
	coll := p.client.Database(os.Getenv("DATABASE")).Collection("prospects")
	err := coll.FindOne(ctx, bson.M{field: value}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
