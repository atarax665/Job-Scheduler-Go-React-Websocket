package jobs

import (
	"scheduler-service/pkg/model"
)

type JobsRepository interface {
	CreateJob(job model.Job) (model.Job, error)
	GetJobs() ([]model.Job, error)
	UpdateJobStatus(status *model.Status) error
}

type jobsRepository struct {
	JobsDataDb map[string]model.Job
}
