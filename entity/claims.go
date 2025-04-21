package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CustomeClaims struct {
	Role   Role          `json:"role"`
	Email  string        `json:"email"`
	UserID bson.ObjectID `json:"userId"`
	jwt.RegisteredClaims
}
