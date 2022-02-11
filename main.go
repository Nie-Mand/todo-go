package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const TODOS_FILE_PATH = "todos.txt"

func main() {
	action, payload := destruct(os.Args)

	content, err := ioutil.ReadFile(TODOS_FILE_PATH)
	if err != nil {
		if action != nil && *action == "init" {
			return
		}
		if !os.IsNotExist(err) {
			fmt.Println("Couldnt open the todo file")
			return
		}
		println()
		fmt.Println("no todos file found, creating one...")
		ioutil.WriteFile(
			TODOS_FILE_PATH,
			[]byte("\n"),
			0644,
		)

	}

	_todos := string(content)
	_todos = strings.Replace(_todos, "\r\n", "\n", -1)
	unfiltered_todos := strings.Split(_todos, "\n")

	// filter empty lines
	todos := []string{}
	for _, todo := range unfiltered_todos {
		if len(todo) > 0 {
			todos = append(todos, todo)
		}
	}

	if action == nil {
		listTodos(todos, true)
	} else {
		switch *action {
		case "init":
			listTodos(todos, false)
			return
		case "add":
			addTodo(strings.Join(payload, " "), todos)
		case "did":
			removeTodo(payload[:1], todos)
		case "clear":
			clearTodo()
		}
	}
}

func destruct(args []string) (action *string, payload []string) {
	if len(args) == 1 {
		return nil, nil
	}
	if len(args) == 2 {
		return &args[1], nil
	}
	return &args[1], args[2:]
}

func listTodos(todos []string, verbose bool) {
	if len(todos) == 0 {
		if verbose {
			fmt.Println("no todos, freedom 💯")
		}
	} else {
		fmt.Println("your todos:")
		for i, todo := range todos {
			fmt.Printf("%d. %s\n", i+1, todo)
		}
	}
}

func addTodo(payload string, todos []string) {
	if payload == "" {
		fmt.Println("🐸 please enter a todo")
		return
	}
	if alreadyExists(payload, todos) {
		fmt.Println("🐸 duh, todo already exists")
		return
	}
	todos = append(todos, strings.ToUpper(payload))
	ioutil.WriteFile(
		TODOS_FILE_PATH,
		[]byte(strings.Join(todos, "\n")),
		0644,
	)
	fmt.Println("todo added 🎉💯, make sure to", payload)
}

func removeTodo(payload []string, todos []string) {
	if len(payload) == 0 {
		fmt.Println("🙄 please enter the ID of the todo you did")
	} else {
		idx, err := strconv.Atoi(payload[0])
		if err != nil {
			fmt.Println("🙁 enter a valid ID")
		} else {
			if idx > len(todos) {
				fmt.Println("🙁 please enter a valid ID")
			} else {
				todos = append(todos[:idx-1], todos[idx:]...)
				ioutil.WriteFile(
					TODOS_FILE_PATH,
					[]byte(strings.Join(todos, "\n")),
					0644,
				)
				fmt.Println("you have todo removed, cool 💯")
			}
		}
	}
}

func clearTodo() {
	ioutil.WriteFile(
		TODOS_FILE_PATH,
		[]byte("\n"),
		0644,
	)
	fmt.Println("you did everythin, cool ig.. 💯")
}

func alreadyExists(payload string, todos []string) bool {
	for _, todo := range todos {
		if todo == strings.ToUpper(payload) {
			return true
		}
	}
	return false
}
