package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/auth"
	"todo/models"
	"todo/todo"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsComplete  bool   `json:"is_complete"`
	UserID      string `json:"user_id"`
}

type Handler struct {
	useCase todo.UseCase
}

func NewHandler(useCase todo.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type createInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	err := h.useCase.Create(c.Request.Context(), inp.Title, inp.Description, user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}

type getResponse struct {
	Tasks []*Task `json:"tasks"`
}

func (h *Handler) Get(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	tasks, err := h.useCase.Get(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getResponse{Tasks: toTasks(tasks)})
}

type deleteInput struct {
	ID string `json:"id"`
}

func (h *Handler) Delete(c *gin.Context) {
	inp := new(deleteInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.Delete(c.Request.Context(), inp.ID, user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type changeInput struct {
	ID     string `json:"id"`
	Status bool   `json:"status"`
}

func (h *Handler) ChangeStatus(c *gin.Context) {
	inp := new(changeInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.ChangeStatus(c.Request.Context(), inp.ID, inp.Status, user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func toTasks(bs []*models.Task) []*Task {
	out := make([]*Task, len(bs))

	for i, b := range bs {
		out[i] = toTask(b)
	}

	return out
}

func toTask(t *models.Task) *Task {
	return &Task{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		IsComplete:  t.IsComplete,
		UserID:      t.UserID,
	}
}
