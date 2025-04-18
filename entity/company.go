package entity

type Company struct {
	Name    string  `json:"name" bson:"name"`
	Logo    string  `json:"logo" bson:"logo"`
	Address Address `json:"address" bson:"address"`
}
