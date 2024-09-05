package datahandler

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/pulsone21/go-todo/internal/entities"
	"github.com/pulsone21/go-todo/internal/logg"
)

// Implement a csv loader
type CSVLoader struct {
	StorePath string
}

func NewCSVLoader(path string) *CSVLoader {
	loader := CSVLoader{
		StorePath: fmt.Sprintf("%s/%s", path, "data.csv"),
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0777)
	}
	if _, err := os.Stat(loader.StorePath); errors.Is(err, os.ErrNotExist) {
		os.Create(loader.StorePath)
	}

	return &loader
}

func (l *CSVLoader) Load() ([]entities.Todo, error) {
	file, err := os.ReadFile(l.StorePath)
	if err != nil {
		return nil, fmt.Errorf("CSVLoader.Load() - error in loading in csv file from %s, %w", l.StorePath, err)
	}

	records, err := csv.NewReader(strings.NewReader(string(file))).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("CSVLoader.Load() - error in reading csv file, %w", err)
	}
	out := []entities.Todo{}

	for _, v := range records {
		t, err := entities.CsvSerialize(v)
		if err != nil {
			logg.Error(err.Error())
		}
		out = append(out, *t)
	}

	return out, nil
}

func (l *CSVLoader) Save(todos []entities.Todo) error {
	var data [][]string

	for _, v := range todos {
		data = append(data, strings.Split(v.ToTable(","), ","))
	}

	file, err := os.Create(l.StorePath)
	if err != nil {
		return fmt.Errorf("error reading or creating file in storage location: %s, %w", l.StorePath, err)
	}

	return csv.NewWriter(file).WriteAll(data)
}
