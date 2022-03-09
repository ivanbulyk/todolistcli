package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ivanbulyk/todolistcli/todo"
	"github.com/spf13/cobra"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone [Task ID, starting from 1]",
	Short: "Mark task as not done",

	Args: cobra.MinimumNArgs(1),

	Run: undoneRun,
}

func init() {
	rootCmd.AddCommand(undoneCmd)

}

func undoneRun(cmd *cobra.Command, args []string) {
	dataFile := "tasks.txt"
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Fatalf("Read items : %v\n", err)
	}
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label", err)
	}
	if i > 0 && i <= len(items) {
		items[i-1].Done = false
		fmt.Printf("%s %v\n", items[i-1].Name, "marked undone")

		err = todo.SaveItems(dataFile, items)
		if err != nil {
			log.Fatalf("Save items : %v\ns", err)
		}
	} else {
		log.Println(i, "doesnt match any item")
	}
}
