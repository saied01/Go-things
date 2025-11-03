package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var tasks []Task
var maxId = 0

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A simple CLI for managing tasks",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Run: func(cmd *cobra.Command, args []string) {
		listTasks(tasks)
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a single task",
	Run: func(cmd *cobra.Command, args []string) {
		addTask(&tasks, args)
	},
}

func Execute() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.Execute()
}

func listTasks(taskList []Task) {
	if len(taskList) == 0 {
		fmt.Println("no tasks yet.")
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"id", "Task", "Created", "done"})

	for _, t := range tasks {
		createdAgo := time.Since(t.Created).Round(time.Minute).String() + " ago"
		row := []string{
			fmt.Sprintf("%d", t.ID),
			t.Content,
			createdAgo,
			fmt.Sprintf("%v", t.Done),
		}
		table.Append(row)
	}
	table.Render()
}

func addTask(taskList *[]Task, taskContent []string) {

	maxId++

	var newTask Task = Task{
		Content: taskContent[0],
		Done:    false,
		Created: time.Now(),
		ID:      maxId,
	}

	*taskList = append(*taskList, newTask)
	fmt.Println("Added task: ", taskContent)
}
