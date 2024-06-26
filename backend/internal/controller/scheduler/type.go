package scheduler

import (
	"scheduler-service/internal/service/status"
	"scheduler-service/pkg/model"
	"scheduler-service/pkg/utils"
	"sync"
)

type jobScheduler struct {
	pq            utils.JobPriorityQueue
	mutex         sync.Mutex
	cond          *sync.Cond
	statusService status.StatusService
}

type JobScheduler interface {
	AddJob(job *model.Job)
	ProcessJobs()
}
