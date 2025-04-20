package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string        `json:"nane" bson:"nane"`
	LastName  string        `json:"last_name" bson:"last_name"`
	Account   Account       `json:"account" bson:"account"`
	Role      Role          `json:"role" bson:"role"`
	Company   Company       `json:"company" bson:"company"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"email"`
}
