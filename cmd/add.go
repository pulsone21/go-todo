package cmd

import (
	"fmt"

	"github.com/pulsone21/go-todo/internal/entities"
	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "add [todo-name]",
	Short: "Adds a new todo with the given name",
	Long:  "Add a a new todo with the given name. Accepts also Flags like -p to select the priority of the todo.",
	Args:  cobra.MinimumNArgs(1),
	Run:   add_todo,
}

func init() {
	rootCmd.AddCommand(add)
	add.Flags().IntP("prio", "p", 0, "The priority of the task, 0 -> 3, Default:0")
}

func add_todo(cmd *cobra.Command, args []string) {
	fmt.Println("i should add a Task")

	name := args[0]
	prio, _ := cmd.Flags().GetInt("prio")
	todo := entities.New(name, len(TODOS)+1, prio)
	TODOS = append(TODOS, *todo)
	HANDLER.Save(TODOS)
}
