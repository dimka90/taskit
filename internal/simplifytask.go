package internal

import "fmt"

type Task interface {
	GetID() string
	GetTitle() string
	GetStatus() string
	MarkComplete() error
	String() string
}
type SimplifyTask struct {
	ID         string
	Title      string
	Iscomplete bool
}

func (t *SimplifyTask) GetID() string {
	return *&t.ID
}
func (t *SimplifyTask) GetTitle() string {
	return *&t.Title
}
func (t *SimplifyTask) GetStatus() string {
	if !t.Iscomplete {
		return "pending"
	}
	return "Done"
}
func (t *SimplifyTask) MarkComplete() error {
	t.Iscomplete = true
	return nil
}

func (t SimplifyTask) String() string {
	return fmt.Sprintf("[%s] [%s] - Status [%s]", t.ID, t.Title, t.GetStatus())

}
