package http

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/auth"
	"todo/models"
	"todo/todo/usecase"
)

func TestCreate(t *testing.T) {
	testUser := &models.User{
		Username: "user",
		Password: "pass",
	}

	r := gin.Default()

	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})
	uc := new(usecase.TodoUseCaseMock)

	RegisterHTTPEndpoints(group, uc)

	inp := &createInput{
		Title:       "testTask",
		Description: "TestDesc",
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("Create", inp.Title, inp.Description, testUser).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/todos", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGet(t *testing.T) {
	testUser := &models.User{
		Username: "user",
		Password: "pass",
	}

	r := gin.Default()

	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})
	uc := new(usecase.TodoUseCaseMock)

	RegisterHTTPEndpoints(group, uc)
	tds := make([]*models.Task, 5)

	for i := 0; i < 5; i++ {
		tds[i] = &models.Task{
			ID:          "id",
			Title:       "title",
			Description: "desc",
			IsComplete:  false,
		}
	}

	uc.On("Get", testUser).Return(tds, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/todos", nil)
	r.ServeHTTP(w, req)

	expectedOut := &getResponse{Tasks: toTasks(tds)}
	expectedOutBody, err := json.Marshal(expectedOut)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedOutBody), w.Body.String())
}
func TestDelete(t *testing.T) {
	testUser := &models.User{
		Username: "user",
		Password: "pass",
	}

	r := gin.Default()

	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})
	uc := new(usecase.TodoUseCaseMock)

	RegisterHTTPEndpoints(group, uc)

	inp := &deleteInput{
		ID: "id",
	}
	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("Delete", inp.ID, testUser).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/todos", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestChangeStatus(t *testing.T) {
	testUser := &models.User{
		Username: "user",
		Password: "pass",
	}

	r := gin.Default()

	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})
	uc := new(usecase.TodoUseCaseMock)

	RegisterHTTPEndpoints(group, uc)
	inp := &changeInput{
		ID:     "id",
		Status: false,
	}
	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("ChangeStatus", inp.ID, inp.Status, testUser).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/todos", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
