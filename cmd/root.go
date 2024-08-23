package cmd

import (
	"fmt"
	"os"

	"github.com/pulsone21/go-todo/internal/datahandler"
	"github.com/pulsone21/go-todo/internal/entities"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	TODOS   []entities.Todo
	HANDLER datahandler.Handler
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "todo",
	Short:   "Versatile todo app under active development",
	Long:    `This todo app is highly configurable with a config file in $HOME/.config/todo/conf.yaml`,
	Aliases: []string{"todo"},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: run_interactive,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	var config Config
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-todo.yaml)")

	homeDir, _ := os.UserHomeDir()
	confFile, err := os.ReadFile(fmt.Sprintf("%s/.config/todo/config.yaml", homeDir))
	if err != nil {
		fmt.Println(fmt.Errorf("error in reading in config file from location: %s, %w", fmt.Sprintf("%s/.config/todo/config.yaml", homeDir), err).Error())
		os.Exit(1)
	}
	err = yaml.Unmarshal(confFile, &config)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	HANDLER, err = config.CreateHandler()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	TODOS, err = HANDLER.Load()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run_interactive(cmd *cobra.Command, args []string) {
	fmt.Println("running interactive.... Lol :D")
}
