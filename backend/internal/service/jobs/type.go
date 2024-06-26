package jobs

import (
	"context"
	"scheduler-service/internal/repository/jobs"
	"scheduler-service/pkg/model"
)

type JobsService interface {
	CreateJob(ctx context.Context, job model.Job) (model.Job, error)
	GetJobs(ctx context.Context) ([]model.Job, error)
	UpdateJobStatus(job *model.Status) error
}

type jobsService struct {
	jobsRepo jobs.JobsRepository
}
