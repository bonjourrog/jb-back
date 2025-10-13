package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/bonjourrog/jb/entity/job"
	"github.com/bonjourrog/jb/repository/application"
	_job "github.com/bonjourrog/jb/repository/job"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type JobService interface {
	NewJob(job job.Post, ctx context.Context) (*job.Post, error)
	GetJobs(filter bson.M, page int, ctx context.Context) ([]job.PostWithCompany, int64, error)
	UpdateJob(job job.Post, ctx context.Context) error
	DeleteJob(job_id bson.ObjectID, user_id bson.ObjectID, ctx context.Context) error
}
type jobService struct {
	jobRepo _job.JobRepository
	appRepo application.ApplicationRepository
}

var ()

func NewPostService(jobRepo _job.JobRepository, appRepo application.ApplicationRepository) JobService {
	return &jobService{
		jobRepo: jobRepo,
		appRepo: appRepo,
	}
}
func (j *jobService) NewJob(job job.Post, ctx context.Context) (*job.Post, error) {

	job.Title = strings.TrimSpace(strings.ToLower(job.Title))
	job.ShortDescription = strings.TrimSpace(job.ShortDescription)
	job.Description = strings.TrimSpace(job.Description)

	if job.Title == "" || job.ShortDescription == "" || job.Description == "" || job.CompanyID == bson.NilObjectID || job.ContractType == "" || job.Industry == "" || job.Schedule == "" {
		return nil, errors.New("some required fields are empty")
	}

	job.ID = bson.NewObjectID()
	job.IsFormalJob = true
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()

	insertedID, err := j.jobRepo.Create(job, ctx)
	if err != nil {
		return nil, err
	}
	job.ID = *insertedID
	return &job, nil
}
func (j *jobService) GetJobs(filter bson.M, page int, ctx context.Context) ([]job.PostWithCompany, int64, error) {
	jobs, total, err := j.jobRepo.GetAll(filter, page, ctx)
	if err != nil {
		return nil, 0, err
	}
	return jobs, total, nil
}
func (j *jobService) UpdateJob(job job.Post, ctx context.Context) error {
	job.Title = strings.TrimSpace(strings.ToLower(job.Title))
	job.ShortDescription = strings.TrimSpace(job.ShortDescription)
	job.Description = strings.TrimSpace(job.Description)

	if job.Title == "" || job.ShortDescription == "" || job.Description == "" || job.CompanyID == bson.NilObjectID || job.ContractType == "" || job.Industry == "" || job.Schedule == "" {
		return errors.New("some required fields are empty")
	}
	job.UpdatedAt = time.Now()
	err := j.jobRepo.Update(job, ctx)
	if err != nil {
		return err
	}
	return nil
}
func (j *jobService) DeleteJob(job_id bson.ObjectID, company_id bson.ObjectID, ctx context.Context) error {
	var (
		appIds []bson.ObjectID
		err    error
	)
	apps, err := j.appRepo.FindByField("job_id", job_id, ctx)
	if err != nil {
		return err
	}
	// currently, company_id field in application is not set when user applies to a job
	// so we cannot verify if the application belongs to the company deleting the job
	// once we set the company_id field in application, we can uncomment the code below
	// for _, app := range *apps {
	//     if app.CompanyID != company_id {
	//         return errors.New("unauthorized: application does not belong to your company")
	//     }
	// }
	if apps != nil && len(*apps) > 0 {
		for _, app := range *apps {
			appIds = append(appIds, app.ID)
		}
	}

	if len(appIds) > 0 {
		err = j.appRepo.DeleteManybyIds(appIds, ctx)
		if err != nil {
			return err
		}
	}

	return j.jobRepo.Delete(job_id, company_id, ctx)
}
