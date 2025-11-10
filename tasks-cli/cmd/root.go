package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var maxId = 0

func getTasksFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dir := filepath.Join(home, ".tasks")
	os.MkdirAll(dir, os.ModePerm)
	return filepath.Join(dir, "tasks.json")
}

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

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove task from list",
	Run: func(cmd *cobra.Command, args []string) {
		removeTask(args)
	},
}

func Execute() {
	// add commands
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(removeCmd)

	// add command flags
	listCmd.Flags().BoolP("all", "a", false, "List all tasks (including completed ones)")

	rootCmd.Execute()
}

func removeTask(task []string) {

	filename := getTasksFilePath()

	var tasks []Task
	fileData, err := os.ReadFile(filename)
	if err == nil && len(fileData) > 0 {
		json.Unmarshal(fileData, &tasks)
	}

	removeId, err := strconv.Atoi(task[0])
	if err != nil {
		fmt.Println("Error converting ID:", err)
	}

	var newTasks []Task
	for _, t := range tasks {
		if t.ID == removeId {
			continue
		}
		newTasks = append(newTasks, t)
	}

	data, err := json.MarshalIndent(newTasks, "", " ")
	if err != nil {
		fmt.Println("Error serializing json:", err)
		return
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Printf("removed task %d", removeId)
}

func completeTask(task []string) {
	filename := getTasksFilePath()

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

	filename := getTasksFilePath()

	file, err := os.Open(filename)
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

	maxId = 0

	filename := getTasksFilePath()

	// Abrir o crear el archivo (modo escritura)
	var tasks []Task
	fileData, err := os.ReadFile(filename)
	if err == nil && len(fileData) > 0 {
		json.Unmarshal(fileData, &tasks)
	}

	for _, t := range tasks {
		if t.ID >= maxId {
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
