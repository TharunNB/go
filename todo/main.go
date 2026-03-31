package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Title string
	Done  bool
}

var filename = "tasks.json"

func loadTasks() []Task {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []Task{}
	}

	var tasks []Task
	json.Unmarshal(data, &tasks)
	return tasks
}

func saveTasks(task []Task) {
	data, _ := json.MarshalIndent(task, "", " ")
	os.WriteFile(filename, data, 0644)
}

func addTask(title string) {
	tasks := loadTasks()
	tasks = append(tasks, Task{Title: title, Done: false})
	saveTasks(tasks)
	fmt.Println("Task added!")
}

func listTasks() {
	tasks := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks yet.")
		return
	}

	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "✓"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Title, status)
	}
}

func doneTask(index int) {
	tasks := loadTasks()

	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid task index")
		return
	}

	tasks[index].Done = true
	saveTasks(tasks)
	fmt.Println("Task marked as completed")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  add \"task name\"")
		fmt.Println("  list")
		fmt.Println("  done <task number>")
		return
	}

	command := os.Args[1]

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Provide the title")
			return
		}
		addTask(os.Args[2])

	case "list":
		listTasks()

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Provide task index")
			return
		}
		index, _ := strconv.Atoi(os.Args[2])
		doneTask(index - 1)

	default:
		fmt.Println("Unknown command")
	}
}
