package status

import (
	"net/http"
	"sync"

	"scheduler-service/internal/service/status"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var controller StatusController
var once sync.Once

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewController() StatusController {
	once.Do(func() {
		statusService := status.NewService()
		controller = &statusController{
			statusService: statusService,
		}
	})
	return controller
}

func InitRoutes(router *gin.Engine) {
	controller := NewController()

	socketRoute := router.
		Group("/")
	{
		socketRoute.GET("/ws", controller.HandleNewConnection)

	}
}
func (q statusController) HandleNewConnection(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.Error(c.Writer, "Could not open websocket connection", http.StatusInternalServerError)
		return
	}
	q.statusService.HandleNewConnection(conn)
}
