package main

import (
	"github.com/dimka90/taskit/internal"
)

func main() {

	task := internal.NewSimplifyTask("Cooking")
	task3 := internal.NewSimplifyTask("Washing Plate")

	task2 := internal.NewPriorityTask(" Run for president", 1)
	taskManager := internal.NewTaskManager()
	taskManager.AddTask(task)
	taskManager.AddTask(task2)
	taskManager.AddTask(task3)
	taskManager.SaveToFile("tasks.json")
	taskManager.ListTask()
}
