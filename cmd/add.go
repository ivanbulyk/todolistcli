package cmd

import (
	"log"

	"github.com/ivanbulyk/todolistcli/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [Task name]",
	Short: "Add task to the list",

	Args: cobra.MinimumNArgs(1),

	Run: addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

}

func addRun(cmd *cobra.Command, args []string) {
	dataFile := "tasks.txt"
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v\n", err)
	}
	for _, x := range args {
		item := todo.Task{Name: x}

		items = append(items, item)
	}

	err = todo.SaveItems(dataFile, items)
	if err != nil {
		log.Printf("%v\n", err)
	}
}
