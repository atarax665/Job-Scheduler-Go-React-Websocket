package model

import (
	"scheduler-service/pkg/enum/job_status"
	"time"
)

type Job struct {
	ID       string               `json:"id"`
	Name     string               `json:"name"`
	Duration time.Duration        `json:"duration"`
	Status   job_status.JobStatus `json:"status"`
}

var JobsDataDb map[string]Job
