package jobs

import (
	"context"

	"scheduler-service/internal/repository/jobs"
	"scheduler-service/pkg/model"
	"sync"
)

var service JobsService
var once sync.Once

func NewService() JobsService {
	once.Do(func() {
		service = &jobsService{
			jobsRepo: jobs.NewRepository(),
		}
	})
	return service
}

func (j *jobsService) CreateJob(ctx context.Context, job model.Job) (model.Job, error) {
	return j.jobsRepo.CreateJob(job)
}

func (j *jobsService) GetJobs(ctx context.Context) ([]model.Job, error) {
	return j.jobsRepo.GetJobs()
}

func (j *jobsService) UpdateJobStatus(status *model.Status) error {
	return j.jobsRepo.UpdateJobStatus(status)
}
