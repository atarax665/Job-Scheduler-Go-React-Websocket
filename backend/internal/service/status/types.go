package status

import (
	"scheduler-service/internal/service/jobs"
	"scheduler-service/pkg/model"
	"sync"

	"github.com/gorilla/websocket"
)

type StatusService interface {
	HandleNewConnection(c *websocket.Conn)
	UpdateJobStatus(job *model.Status)
}

type Client struct {
	conn *websocket.Conn
	send chan *model.Status
}

type statusService struct {
	clients     map[*Client]bool
	broadcast   chan *model.Status
	mutex       *sync.Mutex
	jobsService jobs.JobsService
}
