package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/bonjourrog/jb/entity/job"
	_job "github.com/bonjourrog/jb/repository/job"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type JobService interface {
	NewJob(job job.Post, ctx context.Context) error
	GetJobs(filter bson.M, page int, ctx context.Context) ([]job.PostWithCompany, int64, error)
	UpdateJob(job job.Post, ctx context.Context) error
	DeleteJob(job_id bson.ObjectID, user_id bson.ObjectID, ctx context.Context) error
}
type jobService struct{}

var (
	_jobRepo _job.JobRepository
)

func NewPostService(jobRepo _job.JobRepository) JobService {
	_jobRepo = jobRepo
	return &jobService{}
}
func (*jobService) NewJob(job job.Post, ctx context.Context) error {

	job.Title = strings.TrimSpace(strings.ToLower(job.Title))
	job.ShortDescription = strings.TrimSpace(job.ShortDescription)
	job.Description = strings.TrimSpace(job.Description)

	if job.Title == "" || job.ShortDescription == "" || job.Description == "" || job.CompanyID == bson.NilObjectID || job.ContractType == "" || job.Industry == "" || job.Schedule == "" {
		return errors.New("some required fields are empty")
	}

	job.ID = bson.NewObjectID()
	job.IsFormalJob = true
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()

	if err := _jobRepo.Create(job, ctx); err != nil {
		return err
	}
	return nil
}
func (*jobService) GetJobs(filter bson.M, page int, ctx context.Context) ([]job.PostWithCompany, int64, error) {
	jobs, total, err := _jobRepo.GetAll(filter, page, ctx)
	if err != nil {
		return nil, 0, err
	}
	return jobs, total, nil
}
func (jobService) UpdateJob(job job.Post, ctx context.Context) error {
	job.Title = strings.TrimSpace(strings.ToLower(job.Title))
	job.ShortDescription = strings.TrimSpace(job.ShortDescription)
	job.Description = strings.TrimSpace(job.Description)

	if job.Title == "" || job.ShortDescription == "" || job.Description == "" || job.CompanyID == bson.NilObjectID || job.ContractType == "" || job.Industry == "" || job.Schedule == "" {
		return errors.New("some required fields are empty")
	}
	job.UpdatedAt = time.Now()
	err := _jobRepo.Update(job, ctx)
	if err != nil {
		return err
	}
	return nil
}
func (*jobService) DeleteJob(job_id bson.ObjectID, user_id bson.ObjectID, ctx context.Context) error {
	return _jobRepo.Delete(job_id, user_id, ctx)
}
