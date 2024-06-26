package job_status

type JobStatus string

const (
	Pending   JobStatus = "PENDING"
	Running   JobStatus = "RUNNING"
	Completed JobStatus = "COMPLETED"
)
