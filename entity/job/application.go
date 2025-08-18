package job

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Application struct {
	ID     bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID bson.ObjectID `json:"user_id" bson:"user_id"`
	JobID  bson.ObjectID `json:"job_id" bson:"job_id"`
}
