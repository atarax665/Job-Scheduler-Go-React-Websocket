package scheduler

import (
	"container/heap"
	"fmt"
	"scheduler-service/internal/service/status"
	"scheduler-service/pkg/enum/job_status"
	"scheduler-service/pkg/model"
	"scheduler-service/pkg/utils"
	"sync"
	"time"
)

var scheduler JobScheduler
var once sync.Once

func NewJobScheduler() JobScheduler {
	once.Do(func() {
		pq := make(utils.JobPriorityQueue, 0)
		js := &jobScheduler{
			pq:            pq,
			mutex:         sync.Mutex{},
			statusService: status.NewService(),
		}
		js.cond = sync.NewCond(&js.mutex)
		heap.Init(&pq)
		scheduler = js
	})
	return scheduler
}

func (js *jobScheduler) AddJob(job *model.Job) {
	js.mutex.Lock()
	defer js.mutex.Unlock()
	js.statusService.UpdateJobStatus(&model.Status{
		ID:     job.ID,
		Name:   job.Name,
		Status: job_status.Pending,
	})
	heap.Push(&js.pq, job)
	js.cond.Signal()
}

func (js *jobScheduler) ProcessJobs() {
	for {
		js.mutex.Lock()
		for js.pq.Len() == 0 {
			js.cond.Wait()
		}
		job := heap.Pop(&js.pq).(*model.Job)
		js.mutex.Unlock()
		fmt.Printf("Processing job: %s with duration: %s\n", job.Name, job.Duration)
		time.Sleep(100 * time.Millisecond)
		js.statusService.UpdateJobStatus(&model.Status{
			ID:     job.ID,
			Name:   job.Name,
			Status: job_status.Running,
		})
		time.Sleep(job.Duration)
		js.statusService.UpdateJobStatus(&model.Status{
			ID:     job.ID,
			Name:   job.Name,
			Status: job_status.Completed,
		})
		fmt.Printf("Completed job: %s with duration: %s\n", job.Name, job.Duration)
	}
}
