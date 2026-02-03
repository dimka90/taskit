package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task interface {
	GetID() string
	GetTitle() string
	GetStatus() string
	MarkComplete() error
	String() string
}
type TaskManager struct {
	tasks []Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}
func (tm *TaskManager) AddTask(task Task) {
	tm.tasks = append(tm.tasks, task)
}
func (tm *TaskManager) ListTask() {
	for _, task := range tm.tasks {
		fmt.Println(task)
	}
}

func (tm *TaskManager) CompleteTask(taskId string) (bool, error) {
	for _, task := range tm.tasks {
		if task.GetID() == taskId {
			if err := task.MarkComplete(); err != nil {
				return false, nil
			}
			return true, nil
		}
	}
	return false, fmt.Errorf(" Task not found")
}
func (tm *TaskManager) SaveToFile(filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	json.NewEncoder(file).Encode(tm.tasks)
	return nil
}
