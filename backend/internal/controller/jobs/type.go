package jobs

import (
	"scheduler-service/internal/controller/scheduler"
	"scheduler-service/internal/service/jobs"

	"github.com/gin-gonic/gin"
)

type JobsController interface {
	GetJobs(c *gin.Context)
	CreateJob(c *gin.Context)
}

type jobsController struct {
	jobsService  jobs.JobsService
	jobScheduler scheduler.JobScheduler
}
