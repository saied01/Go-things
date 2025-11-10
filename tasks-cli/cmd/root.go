package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A simple CLI for managing tasks",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Run: func(cmd *cobra.Command, args []string) {
		showAll, _ := cmd.Flags().GetBool("all")
		listTasks(showAll)
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a single task",
	Run: func(cmd *cobra.Command, args []string) {
		addTask(args)
	},
}

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		completeTask(args)
	},
}

func Execute() {
	// add commands
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(completeCmd)

	// add command flags
	listCmd.Flags().BoolP("all", "a", false, "List all tasks (including completed ones)")

	rootCmd.Execute()
}

func completeTask(task []string) {
	filename := "tasks.json"

	var tasks []Task
	fileData, err := os.ReadFile(filename)
	if err == nil && len(fileData) > 0 {
		json.Unmarshal(fileData, &tasks)
	}

	completeId, err := strconv.Atoi(task[0])
	if err != nil {
		fmt.Println("Error converting ID:", err)
	}

	for i, t := range tasks {
		if t.ID == completeId {
			tasks[i].Done = true
			break
		}
	}

	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error serializing json:", err)
		return
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Printf("Marked task %d as completed:", completeId)
}

func listTasks(showAll bool) {

	file, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var tasks []Task
	json.NewDecoder(file).Decode(&tasks)

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"id", "Task", "Created", "done"})

	for _, t := range tasks {
		if !showAll && t.Done {
			continue
		}
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

func addTask(taskContent []string) {

	var maxId int = 0

	filename := "tasks.json"

	// Abrir o crear el archivo (modo escritura)
	var tasks []Task
	fileData, err := os.ReadFile(filename)
	if err == nil && len(fileData) > 0 {
		json.Unmarshal(fileData, &tasks)
	}

	for _, t := range tasks {
		if t.ID > maxId {
			maxId = t.ID + 1
		}
	}

	var newTask Task = Task{
		Content: taskContent[0],
		Done:    false,
		Created: time.Now(),
		ID:      maxId,
	}

	tasks = append(tasks, newTask)

	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error serializing json:", err)
		return
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Added task: ", taskContent[0])
}
