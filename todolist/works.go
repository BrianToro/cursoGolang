package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (tl *taskList) printListOFTasks() {
	for _, task := range tl.tasks {
		fmt.Println(task.name)
		fmt.Println(task.description)
	}
}

func (tl *taskList) printTasksCompleted() {
	for _, task := range tl.tasks {
		if task.complete {
			fmt.Println(task.name)
		}
	}
}

func (tl *taskList) printTaskUncompleted() {
	for _, task := range tl.tasks {
		if !task.complete {
			fmt.Println(task.name)
		}
	}
}

func (tl *taskList) addTask(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) deleteTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

type task struct {
	name        string
	description string
	complete    bool
}

func (t *task) completeTask() {
	t.complete = true
}

func (t *task) updateDescriptionOfTask(newDescription string) {
	t.description = newDescription
}

func (t *task) updateNameOfTask(newName string) {
	t.name = newName
}

func main() {
	t1 := &task{
		name:        "Terminar curso de GO",
		description: "Completar mi curso de go en platzi.com",
		complete:    false,
	}

	t2 := &task{
		name:        "Terminar curso de Nodejs",
		description: "Completar mi curso de nodejs en platzi.com",
		complete:    false,
	}

	t3 := &task{
		name:        "Terminar curso de Python",
		description: "Completar mi curso de python en platzi.com",
		complete:    false,
	}

	t4 := &task{
		name:        "Terminar curso de Java",
		description: "Completar mi curso de java en platzi.com",
		complete:    false,
	}

	listOfTask := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	listOfTask2 := &taskList{tasks: []*task{
		t3, t4,
	}}

	mapOfTasks := make(map[string]*taskList)

	mapOfTasks["Brian"] = listOfTask
	mapOfTasks["Pedrin"] = listOfTask2

	fmt.Println("Task of Brian: ")
	mapOfTasks["Brian"].printListOFTasks()
	fmt.Println("Task of Pedrin: ")
	mapOfTasks["Pedrin"].printListOFTasks()
}
