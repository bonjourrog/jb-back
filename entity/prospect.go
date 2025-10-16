package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Prospect struct {
	ID          bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CompanyName string        `json:"company_name" bson:"company_name"`
	ContactName string        `json:"contact_name" bson:"contact_name"`
	Email       string        `json:"email" bson:"email"`
	Phone       string        `json:"phone" bson:"phone"`
	CompanySize string        `json:"company_size" bson:"company_size"`
	Vacancies   string        `json:"vacancies" bson:"vacancies"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
}
