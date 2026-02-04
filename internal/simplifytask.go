package internal

import "fmt"

type SimplifyTask struct {
	ID         string
	Title      string
	IsComplete bool
}

var id int

func generateID() int {
	id += 1
	return id
}
func NewSimplifyTask(title string) *SimplifyTask {

	return &SimplifyTask{
		ID:         fmt.Sprintf("%d", generateID()),
		Title:      title,
		IsComplete: false,
	}
}
func (t *SimplifyTask) GetID() string {
	return t.ID
}
func (t *SimplifyTask) GetTitle() string {
	return t.Title
}
func (t *SimplifyTask) GetStatus() string {
	if !t.IsComplete {
		return "pending"
	}
	return "Done"
}
func (t *SimplifyTask) MarkComplete() error {
	t.IsComplete = true
	return nil
}

func (t SimplifyTask) String() string {
	return fmt.Sprintf("[%s] [%s] - Status [%s]", t.ID, t.Title, t.GetStatus())

}
