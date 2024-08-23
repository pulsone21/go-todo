package datahandler

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pulsone21/go-todo/internal/entities"
)

// Implement a csv loader
type CSVLoader struct {
	StorePath string
}

func NewCSVLoader(path string) *CSVLoader {
	loader := CSVLoader{
		StorePath: fmt.Sprintf("%s/%s", path, "data.csv"),
	}

	if _, err := os.Stat(loader.StorePath); errors.Is(err, os.ErrNotExist) {
		os.Create(loader.StorePath)
	}

	return &loader
}

func (l *CSVLoader) Load() ([]entities.Todo, error) {
	file, err := os.ReadFile(l.StorePath)
	if err != nil {
		return nil, fmt.Errorf("error in loading in csv file from %s, %w", l.StorePath, err)
	}

	records, err := csv.NewReader(strings.NewReader(string(file))).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error in reading csv file, %w", err)
	}
	out := []entities.Todo{}
	for _, v := range records {

		prio, err := strconv.Atoi(v[1])
		if err != nil {
			return nil, fmt.Errorf("error in parsing priority string to int, %w", err)
		}

		state, err := strconv.Atoi(v[2])
		if err != nil {
			return nil, fmt.Errorf("error in parsing state string to int, %w", err)
		}

		doneTS, err := strconv.Atoi(v[3])
		if err != nil {
			return nil, fmt.Errorf("error in parsing doneTS string to int, %w", err)
		}
		out = append(out, entities.Todo{
			Name:     v[0],
			Priority: prio,
			State:    entities.TodoState(state),
			DoneTS:   int64(doneTS),
		})
	}

	return out, nil
}

func (l *CSVLoader) Save(todos []entities.Todo) error {
	var data [][]string

	for _, v := range todos {
		csvStr := fmt.Sprintf("%s,%i,%s,%i\n", v.Name, v.Priority, v.State, v.DoneTS)
		data = append(data, strings.Split(csvStr, ","))
	}

	file, err := os.Create(l.StorePath)
	if err != nil {
		return fmt.Errorf("error reading or creating file in storage location: %s, %w", l.StorePath, err)
	}

	return csv.NewWriter(file).WriteAll(data)
}
