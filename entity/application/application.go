package application

import (
	"time"

	"github.com/bonjourrog/jb/entity/job"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Application struct {
	ID        bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    bson.ObjectID `json:"user_id" bson:"user_id"`
	CompanyID bson.ObjectID `json:"company_id" bson:"company_id"`
	JobID     bson.ObjectID `json:"job_id" bson:"job_id"`
	Status    Status        `json:"status" bson:"status"`
	AppliedAt time.Time     `json:"applied_at" bson:"applied_at"`
}

type ApplicationWithCompany struct {
	Application         `json:",inline" bson:",inline"`
	job.PostWithCompany `json:",inline" bson:",inline"`
}
