package cmd

import (
	"fmt"

	"github.com/pulsone21/go-todo/internal/entities"
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "Adds a new todo with the given name",
	Long:  "Add a a new todo with the given name. Accepts also Flags like -p to select the priority of the todo.",
	Run:   list_todos,
}

func init() {
	rootCmd.AddCommand(list)
	list.Flags().IntP("prio", "p", -1, "The priority of the task, 0 -> 3, Default:-1 (shows all)")
	list.Flags().BoolP("done", "d", false, "If done Task should be shown, ignored if state flag is set")
	list.Flags().IntP("state", "s", -1, "Filter by state, Default: -1 Shows all")
	list.MarkFlagsMutuallyExclusive("done", "state")
}

func list_todos(cmd *cobra.Command, args []string) {
	// prio, _ := cmd.Flags().GetInt("prio")
	// state, _ := cmd.Flags().GetInt("state")
	// done, _ := cmd.Flags().GetBool("done")
	//
	// to_show := []entities.Todo{}

	if len(TODOS) == 0 {
		fmt.Println("All todos done nice 🎉")
		return
	}

	Print(entities.TableHeader())
	for _, v := range TODOS {
		Print(v.TableFormat())
	}
}
