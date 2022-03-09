package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/ivanbulyk/todolistcli/todo"
	"github.com/spf13/cobra"
)

var (
	doneOpt bool
	allOpt  bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks still to do",

	Run: listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "undone", false, "Show 'done' todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all todos")
}

func listRun(cmd *cobra.Command, args []string) {
	dataFile := "tasks.txt"
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v\n", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		if allOpt || i.Done == doneOpt {
			fmt.Fprintln(w, i.Label()+"\t"+"\t"+i.Name+"\t")
		}
	}
	w.Flush()
}
