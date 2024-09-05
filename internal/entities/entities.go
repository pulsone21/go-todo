package entities

import (
	"fmt"
	"reflect"
	"strconv"
)

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
	return fmt.Sprintf("%b,%s,%s,%s,%b", t.Id, t.Name, t.PrioToString(), t.State.ToString(), t.DoneTS)
}

func (t *Todo) ToTable(separator string) string {
	v := reflect.ValueOf(t)
	h := v.Type()
	result := ""

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		field := h.Field(i)
		fixStr := ""

		switch fmt.Sprint(field) {
		case "Priority":
			fixStr = t.PrioToString()
		case "State":
			fixStr = value.(*TodoState).ToString()
		default:
			fixStr = value.(string)
		}

		result += fmt.Sprintf("%v%s", fixStr, separator)
	}
	return result
}

func CsvSerialize(values []string) (*Todo, error) {
	id, err := strconv.Atoi(values[0])
	if err != nil {
		return nil, fmt.Errorf("CSVLoader.Load() - error in parsing id string to int, %w", err)
	}

	prio, err := strconv.Atoi(values[2])
	if err != nil {
		return nil, fmt.Errorf("CSVLoader.Load() - error in parsing priority string to int, %w", err)
	}

	state, err := strconv.Atoi(values[3])
	if err != nil {
		return nil, fmt.Errorf("CSVLoader.Load() - error in parsing state string to int, %w", err)
	}

	doneTS, err := strconv.Atoi(values[4])
	if err != nil {
		return nil, fmt.Errorf("CSVLoader.Load() - error in parsing doneTS string to int, %w", err)
	}

	return &Todo{
		Id:       id,
		Name:     values[1],
		Priority: prio,
		State:    TodoState(state),
		DoneTS:   int64(doneTS),
	}, nil
}

func GenerateHeader(todo Todo, separator string) string {
	v := reflect.ValueOf(todo)
	t := v.Type()
	result := ""

	for i := 0; i < v.NumField(); i++ {
		result += fmt.Sprintf("%v%s", t.Field(i), separator)
	}
	return result
}

func (tS *TodoState) ToString() string {
	switch *tS {
	case Open:
		return "Open"
	case Doing:
		return "Doing"
	case Done:
		return "Done"
	default:
		return fmt.Sprintf("Unkown - %b", tS)
	}
}

func (t *Todo) PrioToString() string {
	strings := []string{"low", "medium", "high"}
	if t.Priority > len(strings) {
		return fmt.Sprintf("Unknown priority: %b", t.Priority)
	}
	return strings[t.Priority]
}
