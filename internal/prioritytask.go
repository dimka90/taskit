package internal

import "fmt"

type PriorityTask struct {
	ID         string
	Title      string
	Iscomplete bool
	Priority   int
}

var taskID int

func NewPriorityTask(title string, priority int) *PriorityTask {
	return &PriorityTask{
		ID:         fmt.Sprintf("%d", generateID()),
		Title:      title,
		Iscomplete: false,
		Priority:   priority,
	}

}
func (t *PriorityTask) GetID() string {
	return t.ID
}

func (t *PriorityTask) GetTitle() string {
	return t.Title
}
func (t *PriorityTask) GetStatus() string {
	if !t.Iscomplete {
		return "Pending"
	}
	return "Done"
}

func (t *PriorityTask) MarkComplete() error {
	*&t.Iscomplete = true
	return nil
}
func (t PriorityTask) String() string {
	return fmt.Sprintf("[%s] - [%s] [%s] - priority [%d]", t.ID, t.Title, t.GetStatus(), t.Priority)
}
