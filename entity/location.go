package entity

type Location struct {
	Type        string    `json:"type" bson:"type"`               // always "Point"
	Coordinates []float64 `json:"coordinates" bson:"coordinates"` // [longitude, latitude]
}
