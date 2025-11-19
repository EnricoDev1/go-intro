package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/fatih/color"
)

type Task struct {
	id      int
	title   string
	content string
}

var tasks = []Task{}

func menu() {
	color.Cyan("Insert choice")
	color.Cyan("1. Add task")
	color.Cyan("2. Delete task")
	color.Cyan("3. View tasks")
	color.Cyan("4. Exit")
	color.RGB(200, 200, 42).Print("Choice: ")
}

func addTask() {
	var task Task

	fmt.Print("Insert title: ")
	fmt.Scanln(&task.title)

	fmt.Print("Insert title: ")
	fmt.Scanln(&task.content)

	task.id = rand.IntN(99999999)

	tasks = append(tasks, task)
}

func viewTasks() {
	if len(tasks) == 0 {
		color.Red("Nothing to show")
		return
	}

	for i := 0; i < len(tasks); i++ {
		fmt.Println("-----------------------")
		fmt.Printf("Id: %d\n", tasks[i].id)
		fmt.Printf("Title: %s\n", tasks[i].title)
		fmt.Printf("content: %s\n", tasks[i].content)
	}
	fmt.Println("-----------------------")
}

func deleteTask() {
	var id int
	fmt.Println("Insert the id of the task to remove: ")
	fmt.Scanln(&id)

	for i := 0; i < len(tasks); i++ {
		if tasks[i].id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
}

func main() {
	var choice int = 0
	for choice != 4 {
		menu()
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask()
		case 2:
			deleteTask()
		case 3:
			viewTasks()
		case 4:
			fmt.Println("bye")
		}
	}
}
