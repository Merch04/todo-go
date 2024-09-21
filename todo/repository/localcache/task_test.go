package localcache

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"todo/models"
	"todo/todo"
)

func TestGetTask(t *testing.T) {
	s := NewTaskLocalStorage()

	idUser := "iduser"
	idUser2 := "iduser2"
	idTask := "idtask"
	user := &models.User{ID: idUser}
	user2 := &models.User{ID: idUser2}
	task := &models.Task{
		ID:          idTask,
		Title:       "Уборка",
		Description: "Ну там сям метануть, пыль протереть, короче ты пон",
		UserID:      user.ID,
	}
	//create
	err := s.CreateTask(context.Background(), task)
	assert.NoError(t, err)

	//get
	//Получаем хозяином
	returnedTasks, err := s.GetTasks(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, task, returnedTasks[0])
	//получаем нн-ом
	returnedTasks, err = s.GetTasks(context.Background(), user2)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(returnedTasks))

}
func TestDeleteTask(t *testing.T) {
	s := NewTaskLocalStorage()

	idUser := "iduser"
	idUser2 := "iduser2"
	idTask := "idtask"
	user := &models.User{ID: idUser}
	user2 := &models.User{ID: idUser2}
	task := &models.Task{
		ID:          idTask,
		Title:       "Уборка",
		Description: "Ну там сям метануть, пыль протереть, короче ты пон",
		UserID:      user.ID,
	}

	//create
	err := s.CreateTask(context.Background(), task)
	assert.NoError(t, err)

	//get
	returnedTasks, err := s.GetTasks(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, task, returnedTasks[0])

	//delete
	//удаляем нн-ом
	err = s.DeleteTask(context.Background(), idTask, user2)
	assert.Error(t, err)
	assert.Equal(t, err, todo.ErrTaskNotFound)

	returnedTasks, err = s.GetTasks(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(returnedTasks))

	//удаляем хозяином
	err = s.DeleteTask(context.Background(), idTask, user)
	assert.NoError(t, err)

	returnedTasks, err = s.GetTasks(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(returnedTasks))
}
func TestChangeTask(t *testing.T) {
	s := NewTaskLocalStorage()

	idUser := "iduser"
	idUser2 := "iduser2"
	idTask := "idtask"
	user := &models.User{ID: idUser}
	user2 := &models.User{ID: idUser2}
	task := &models.Task{
		ID:          idTask,
		Title:       "Уборка",
		Description: "Ну там сям метануть, пыль протереть, короче ты пон",
		UserID:      user.ID,
	}

	//create
	err := s.CreateTask(context.Background(), task)
	assert.NoError(t, err)

	//get
	returnedTasks, err := s.GetTasks(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, task, returnedTasks[0])

	//change
	//меняем хозяином
	err = s.ChangeStatus(context.Background(), idTask, true, user)
	assert.NoError(t, err)
	returnedTasks, err = s.GetTasks(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, true, returnedTasks[0].IsComplete)
	//меняем нн-ом
	err = s.ChangeStatus(context.Background(), idTask, false, user2)
	assert.Error(t, err)
	returnedTasks, err = s.GetTasks(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, true, returnedTasks[0].IsComplete)
}
