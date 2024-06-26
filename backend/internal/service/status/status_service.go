package status

import (
	"encoding/json"
	"log"
	"scheduler-service/internal/service/jobs"
	"scheduler-service/pkg/model"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	service StatusService
	once    sync.Once
)

func NewService() StatusService {
	once.Do(func() {
		service = &statusService{
			clients:     make(map[*Client]bool),
			broadcast:   make(chan *model.Status),
			mutex:       &sync.Mutex{},
			jobsService: jobs.NewService(),
		}
	})
	return service
}

func (cs statusService) HandleNewConnection(conn *websocket.Conn) {
	client := &Client{conn: conn, send: make(chan *model.Status)}
	cs.mutex.Lock()
	cs.clients[client] = true
	cs.mutex.Unlock()
	go handleWrites(client)
	go cs.broadcastMessages()

}

func handleWrites(client *Client) {
	// defer client.conn.Close()

	for msg := range client.send {
		data, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Error marshaling message: %v", err)
			return
		}

		err = client.conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Printf("Error writing message: %v", err)
			return
		}
	}
}

func (cs statusService) broadcastMessages() {
	for {
		msg := <-cs.broadcast
		if len(cs.clients) == 0 {
			continue
		}
		cs.mutex.Lock()
		for client := range cs.clients {
			select {
			case client.send <- msg:
			default:
				close(client.send)
				delete(cs.clients, client)
			}
		}
		cs.mutex.Unlock()
	}
}

func (cs statusService) UpdateJobStatus(job *model.Status) {
	msg := job
	cs.jobsService.UpdateJobStatus(job)
	cs.broadcast <- msg
	log.Println("Job status updated: ", job.Status)
}
