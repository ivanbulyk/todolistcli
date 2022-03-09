package cmd

import (
	"log"

	"github.com/ivanbulyk/todolistcli/todo"
	"github.com/spf13/cobra"
)

var (
	alreadyDoneOpt bool
	allDoneOpt     bool
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup done tasks",

	Run: cleanupRun,
}

func init() {
	rootCmd.AddCommand(cleanupCmd)

	listCmd.Flags().BoolVar(&alreadyDoneOpt, "done", true, "Clean 'done' todos")
	listCmd.Flags().BoolVar(&allDoneOpt, "all done", true, "Clean all todos")
}

func cleanupRun(cmd *cobra.Command, args []string) {
	dataFile := "tasks.txt"
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v\n", err)
	}

	for index, i := range items {

		if i.Done == alreadyDoneOpt {

			copy(items[index:], items[index+1:]) // Shift a[i+1:] left one index.
			items[len(items)-1] = todo.Task{}    // Erase last element (write zero value).
			items = items[:len(items)-1]         // Truncate slice.
		}
	}

	err = todo.SaveItems(dataFile, items)
	if err != nil {
		log.Printf("%v\n", err)
	}
}
