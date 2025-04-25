package job

import (
	"time"

	"github.com/bonjourrog/jb/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Post struct {
	ID               bson.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Title            string          `json:"title" bson:"title"`
	ShortDescription string          `json:"short_description" bson:"short_description"`
	Description      string          `json:"description" bson:"description"`
	Salary           float32         `json:"salary" bson:"salary"`
	Benefits         []string        `json:"benefits" bson:"benefits"`
	Location         entity.Location `json:"location" bson:"location"`
	Industry         string          `json:"industry" bson:"industry"`
	Schedule         string          `json:"schedule" bson:"schedule"`
	ContractType     string          `json:"contract_type" bson:"contract_type"`
	IsFormalJob      bool            `json:"is_formal_job" bson:"location"`
	Published        bool            `json:"published" bson:"published"`
	CompanyID        string          `json:"company_id" bson:"company_id"`
	CreatedAt        time.Time       `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" bson:"updated_at"`
}
