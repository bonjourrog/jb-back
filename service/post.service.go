package service

import (
	"errors"
	"strings"
	"time"

	"github.com/bonjourrog/jb/entity/job"
	_job "github.com/bonjourrog/jb/repository/job"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type JobService interface {
	NewJob(job job.Post) error
	GetJobs(filter bson.M, page int) ([]job.PostWithCompany, int64, error)
	UpdateJob(job job.Post) error
	DeleteJob(job_id bson.ObjectID, user_id bson.ObjectID) error
	ApplyToJob(user_id string, job_id string) error
}
type jobService struct{}

var (
	_jobRepo _job.JobRepository
)

func NewPostService(jobRepo _job.JobRepository) JobService {
	_jobRepo = jobRepo
	return &jobService{}
}
func (*jobService) NewJob(job job.Post) error {

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

	if err := _jobRepo.Create(job); err != nil {
		return err
	}
	return nil
}
func (*jobService) GetJobs(filter bson.M, page int) ([]job.PostWithCompany, int64, error) {
	jobs, total, err := _jobRepo.GetAll(filter, page)
	if err != nil {
		return nil, 0, err
	}
	if len(jobs) == 0 {
		return nil, 0, errors.New("no jobs found")
	}
	return jobs, total, nil
}
func (jobService) UpdateJob(job job.Post) error {
	job.Title = strings.TrimSpace(strings.ToLower(job.Title))
	job.ShortDescription = strings.TrimSpace(job.ShortDescription)
	job.Description = strings.TrimSpace(job.Description)

	if job.Title == "" || job.ShortDescription == "" || job.Description == "" || job.CompanyID == bson.NilObjectID || job.ContractType == "" || job.Industry == "" || job.Schedule == "" {
		return errors.New("some required fields are empty")
	}
	job.UpdatedAt = time.Now()
	err := _jobRepo.Update(job)
	if err != nil {
		return err
	}
	return nil
}
func (*jobService) DeleteJob(job_id bson.ObjectID, user_id bson.ObjectID) error {
	return _jobRepo.Delete(job_id, user_id)
}
func (*jobService) ApplyToJob(user_id string, job_id string) error {
	var (
		application job.Application
	)
	application.ID = bson.NewObjectID()
	UserID, err := bson.ObjectIDFromHex(user_id)
	if err != nil {
		return err
	}
	application.UserID = UserID
	JobID, err := bson.ObjectIDFromHex(job_id)
	if err != nil {
		return err
	}
	application.JobID = JobID
	if user_id == "" || job_id == "" {
		return errors.New("user id or job id is empty")
	}

	alreadyApplied, err := _jobRepo.IsUserAlreadyApplied(UserID, JobID)
	if err != nil {
		return err
	}
	if alreadyApplied {
		return errors.New("user has already applied to this job")
	}

	application.AppliedAt = time.Now()

	if err := _jobRepo.ApplyToJob(application); err != nil {
		return err
	}
	return nil
}
