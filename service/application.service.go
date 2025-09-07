package service

import (
	"context"

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
}

type applicationService struct {
	appRepo repo.ApplicationRepository
	jobRepo jobRepo.JobRepository
}

func NewApplicationService(appRepo repo.ApplicationRepository, jobRepo jobRepo.JobRepository) ApplicationService {
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
func (*applicationService) UpdateStatus(application_id bson.ObjectID, status string, ctx context.Context) error {
	return nil
}
