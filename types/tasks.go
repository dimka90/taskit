package types

type SimplifyTask struct {
	ID         string
	Title      string
	Iscomplete bool
}

type PriorityTask struct {
	id         int
	title      string
	Iscomplete bool
	priority   int
}
