package model

import "scheduler-service/pkg/enum/job_status"

type Status struct {
	ID     string               `json:"id"`
	Name   string               `json:"name"`
	Status job_status.JobStatus `json:"status"`
}
