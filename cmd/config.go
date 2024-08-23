package cmd

import (
	"fmt"
	"strings"

	"github.com/pulsone21/go-todo/internal/datahandler"
)

type Config struct {
	DataHandler string `yaml:"data_handler"`
	StoragePath string `yaml:"storage_path"`
	minWidth    int    `yaml:"min_width"`
	tabWidth    int    `yaml:"tab_width"`
	padChar     byte   `yaml:"pad_char"`
	padding     int    `yaml:"padding"`
}

func (c Config) CreateHandler() (datahandler.Handler, error) {
	switch strings.ToUpper(c.DataHandler) {
	case "CSV":
		return datahandler.NewCSVLoader(c.StoragePath), nil
	default:
		return nil, fmt.Errorf("error in datahandler creation, %s is not implemented.", c.DataHandler)
	}
}
