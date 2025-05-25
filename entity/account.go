package entity

type Account struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Phone    string `json:"phone" bson:"phone"`
	Banned   bool   `json:"banned" bson:"banned"`
}
