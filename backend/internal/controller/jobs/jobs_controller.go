package jobs

import (
	"context"
	"net/http"
	"scheduler-service/internal/controller/scheduler"
	"scheduler-service/internal/service/jobs"
	"scheduler-service/pkg/model"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var controller JobsController
var once sync.Once

func NewController() JobsController {
	once.Do(func() {
		jobScheduler := scheduler.NewJobScheduler()
		go jobScheduler.ProcessJobs()
		controller = &jobsController{
			jobsService:  jobs.NewService(),
			jobScheduler: jobScheduler,
		}
	})
	return controller
}

func InitRoutes(router *gin.Engine) {
	controller := NewController()

	jobsRoutes := router.
		Group("/")
	{
		jobsRoutes.GET("/jobs", controller.GetJobs)
		jobsRoutes.POST("/jobs", controller.CreateJob)
	}
}

func (jc *jobsController) GetJobs(c *gin.Context) {
	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	jobsDetails, err := jc.jobsService.GetJobs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobsDetails)
}

func (jc *jobsController) CreateJob(c *gin.Context) {
	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var jobDetails model.Job
	if err := c.ShouldBindJSON(&jobDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobDetails.Duration = jobDetails.Duration * time.Second
	id := uuid.New().String()
	jobDetails.ID = id

	jobDetail, err := jc.jobsService.CreateJob(c.Request.Context(), jobDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jc.jobScheduler.AddJob(&jobDetail)

	c.JSON(http.StatusOK, jobDetail)
}
