package entities

import "fmt"

type Todo struct {
	Id       int
	Name     string
	Priority int
	State    TodoState
	DoneTS   int64
}

type TodoState int

const (
	Open  TodoState = 0
	Doing TodoState = 1
	Done  TodoState = 2
)

func New(name string, id, prio int) *Todo {
	return &Todo{
		Id:       id,
		Name:     name,
		Priority: prio,
		State:    Open,
		DoneTS:   -1,
	}
}

func (t *Todo) TableFormat() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s", t, t.Id, t.Name, t.Priority, t.State, t.DoneTS)
}

func TableHeader() string {
	return fmt.Sprintf("ID\tName\tPriority\tState\tDone Time")
}
