package entity

type Account struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Banned   bool   `json:"banned" bson:"banned"`
}
