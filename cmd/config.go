package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pulsone21/go-todo/internal/datahandler"
	"github.com/pulsone21/go-todo/internal/logg"
)

type Config struct {
	DataHandler string `yaml:"data_handler"`
	StoragePath string `yaml:"storage_path"`
	MinWidth    int    `yaml:"min_width"`
	TabWidth    int    `yaml:"tab_width"`
	PadChar     string `yaml:"pad_char"`
	Padding     int    `yaml:"padding"`
	LogPath     string `yaml:"log_path"`
	LogLevel    int    `yaml:"log_level"`
	LogSize     int    `yaml:"log_size"`
}

func NewConfig() Config {
	home_dir := os.Getenv("$HOME")
	return Config{
		DataHandler: "CSV",
		StoragePath: fmt.Sprintf("%v/.config/todo/data", home_dir),
		MinWidth:    2,
		TabWidth:    4,
		PadChar:     " ",
		Padding:     4,
		LogPath:     fmt.Sprintf("%v/.config/todo", home_dir),
		LogLevel:    0,
		LogSize:     20,
	}
}

func (c *Config) CreateHandler() (datahandler.Handler, error) {
	path := ensure_path(c.StoragePath)
	logg.Info(fmt.Sprintf("ensured path: %v\n", path))
	switch strings.ToUpper(c.DataHandler) {
	case "CSV":
		return datahandler.NewCSVLoader(path), nil
	default:
		return nil, fmt.Errorf("error in datahandler creation, %s is not implemented.", c.DataHandler)
	}
}

func ensure_path(path string) string {
	str_le := len(path)
	for strings.HasSuffix(path, "/") {
		path = path[:str_le-1]
	}
	return path
}
