package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


const fileName = "todo.txt"

func AddTodo(task string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("[ ] " + task + "\n")
	return err
}

func ListTodos() error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 1
	for scanner.Scan() {
		fmt.Printf("%d. %s\n", i, scanner.Text())
		i++
	}
	return scanner.Err()
}


func MarkDone(index int) error {
	f, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	lines := strings.Split(string(f), "\n")

	if index < 1 || index > len(lines) {
		return fmt.Errorf("task number out of range")
	}

	if lines[index-1] == "" {
		return fmt.Errorf("no such task")
	}

	if strings.HasPrefix(lines[index-1], "[x]") {
		fmt.Println("Task already marked as done.")
		return nil
	}

	lines[index-1] = strings.Replace(lines[index-1], "[ ]", "[x]", 1)

	return os.WriteFile(fileName, []byte(strings.Join(lines, "\n")), 0644)
}