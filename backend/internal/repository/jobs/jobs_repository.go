package jobs

import (
	"scheduler-service/pkg/enum/job_status"
	"scheduler-service/pkg/message"
	"scheduler-service/pkg/model"
	"sync"
)

var jobsRepo JobsRepository
var once sync.Once

func NewRepository() JobsRepository {
	once.Do(func() {
		jobsRepo = &jobsRepository{
			JobsDataDb: make(map[string]model.Job),
		}
	})
	return jobsRepo
}

func (j *jobsRepository) CreateJob(job model.Job) (model.Job, error) {
	job.Status = job_status.Pending
	j.JobsDataDb[job.ID] = job
	return job, nil
}

func (j *jobsRepository) GetJobs() ([]model.Job, error) {
	var jobs []model.Job
	for _, job := range j.JobsDataDb {
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (j *jobsRepository) UpdateJobStatus(status *model.Status) error {
	job, ok := j.JobsDataDb[status.ID]
	if !ok {
		return message.ErrJobNotFound
	}
	job.Status = status.Status
	j.JobsDataDb[status.ID] = job
	return nil
}
