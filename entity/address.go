package entity

type Address struct {
	FirstStreet  string   `json:"first_street" bson:"first_street"`
	SecondStreet string   `json:"second_street" bson:"second_street"`
	Neighborhood string   `json:"neighborhood" bson:"neighborhood"`
	Location     Location `json:"location" bson:"location"`
}
