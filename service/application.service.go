package service

import (
	"context"
	"errors"
	"time"

	"github.com/bonjourrog/jb/entity/application"
	"github.com/bonjourrog/jb/entity/job"
	repo "github.com/bonjourrog/jb/repository/application"
	jobRepo "github.com/bonjourrog/jb/repository/job"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ApplicationService interface {
	// Create(application application, ctx context.Context) error
	GetUserApplications(user_id bson.ObjectID, ctx context.Context) ([]application.ApplicationWithCompany, error)
	UpdateStatus(application_id bson.ObjectID, status string, ctx context.Context) error
	ApplyToJob(user_id string, job_id string, ctx context.Context) error
}

type applicationService struct {
	appRepo repo.ApplicationRepository
	jobRepo jobRepo.JobRepository
}

func NewApplicationService(appRepo repo.ApplicationRepository) ApplicationService {
	return &applicationService{
		appRepo: appRepo,
	}
}

func (a *applicationService) GetUserApplications(user_id bson.ObjectID, ctx context.Context) ([]application.ApplicationWithCompany, error) {
	var (
		applications []application.ApplicationWithCompany
		err          error
	)
	apps, err := a.appRepo.FindByUser(user_id, ctx)
	if err != nil {
		return nil, err
	}
	if len(apps) == 0 {
		return []application.ApplicationWithCompany{}, nil
	}

	jobsIds := make([]bson.ObjectID, 0, len(apps))
	for _, app := range apps {
		jobsIds = append(jobsIds, app.JobID)
	}

	// jobs, err := a.jobRepo.GetByIds(jobsIds, ctx)
	jobs, err := a.appRepo.GetByIds(jobsIds, ctx)
	if err != nil {
		return nil, err
	}

	jobMap := make(map[bson.ObjectID]job.PostWithCompany)
	for _, j := range jobs {
		jobMap[j.ID] = j
	}

	for _, app := range apps {
		if j, ok := jobMap[app.JobID]; ok {
			applications = append(applications, application.ApplicationWithCompany{
				app,
				j,
			})
		}
	}
	return applications, nil
}
func (a *applicationService) UpdateStatus(application_id bson.ObjectID, status string, ctx context.Context) error {
	return nil
}
func (a *applicationService) ApplyToJob(user_id string, job_id string, ctx context.Context) error {
	var (
		app application.Application
	)
	app.ID = bson.NewObjectID()
	UserID, err := bson.ObjectIDFromHex(user_id)
	if err != nil {
		return err
	}
	app.UserID = UserID
	JobID, err := bson.ObjectIDFromHex(job_id)
	if err != nil {
		return err
	}
	app.JobID = JobID
	if user_id == "" || job_id == "" {
		return errors.New("user id or job id is empty")
	}

	app.Status = application.StatusReceived

	alreadyApplied, err := a.appRepo.IsUserAlreadyApplied(UserID, JobID, ctx)
	if err != nil {
		return err
	}
	if alreadyApplied {
		return errors.New("user has already applied to this job")
	}

	app.AppliedAt = time.Now()

	if err := a.appRepo.ApplyToJob(app, ctx); err != nil {
		return err
	}
	return nil
}
