package internal

type Task interface {
	GetID() string
	GetTitle() string
	GetStatus() string
	MarkComplete() error
	String() string
}
