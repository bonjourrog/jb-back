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
	IsFormalJob      bool            `json:"is_formal_job" bson:"is_formal_job"`
	Published        bool            `json:"published" bson:"published"`
	CompanyID        bson.ObjectID   `json:"company_id,omitempty" bson:"company_id,omitempty"`
	Slug             string          `json:"slug" bson:"slug"`
	CreatedAt        time.Time       `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" bson:"updated_at"`
}

type PostWithCompany struct {
	Post         `bson:",inline"`
	CompanyName  string `bson:"company_name" json:"company_name"`
	CompanyLogo  string `bson:"company_logo" json:"company_logo"`
	CompanyPhone string `bson:"company_phone" json:"company_phone"`
}
