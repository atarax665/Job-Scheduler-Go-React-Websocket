package jobs_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"scheduler-service/internal/controller/jobs"
	"scheduler-service/internal/controller/status"
	"scheduler-service/internal/service/jobs/mocks"
	"scheduler-service/pkg/model"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	jobs.InitRoutes(router)
	status.InitRoutes(router)
	return router
}

func TestCreateJobSuccess(t *testing.T) {
	mockJobsService := new(mocks.JobsService)
	mockJobScheduler := new(mocks.JobScheduler)

	router := setupRouter()
	connectToWS(t, httptest.NewServer(router))

	jobDetails := model.Job{
		Name:     "Test Job",
		Duration: 10,
	}
	jobDetailsJSON, _ := json.Marshal(jobDetails)

	mockJobsService.On("CreateJob", mock.Anything, mock.AnythingOfType("model.Job")).Return(jobDetails, nil)
	mockJobScheduler.On("AddJob", mock.AnythingOfType("*model.Job")).Return()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/jobs", bytes.NewBuffer(jobDetailsJSON))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response model.Job
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Test Job", response.Name)
	assert.Equal(t, 10*time.Second, response.Duration)
}

func TestCreateJobBindJSONError(t *testing.T) {
	router := setupRouter()

	invalidJSON := []byte(`{invalid json}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/jobs", bytes.NewBuffer(invalidJSON))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetJobsSuccess(t *testing.T) {
	mockJobsService := new(mocks.JobsService)

	router := setupRouter()
	connectToWS(t, httptest.NewServer(router))
	jobDetails := model.Job{
		Name:     "Test Job",
		Duration: 10,
	}

	mockJobsService.On("GetJobs", mock.Anything).Return(jobDetails, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/jobs", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response []model.Job
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, "Test Job", response[0].Name)
	assert.Equal(t, 10*time.Second, response[0].Duration)
}

func connectToWS(t *testing.T, server *httptest.Server) *websocket.Conn {
	wsURL := "ws" + server.URL[len("http"):] + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.NoError(t, err)
	return conn
}
