package main

import (
		"fmt" 
		"os"
		"strconv"
	)

func main(){
	
	if len(os.Args)< 2{
		fmt.Println("Usage: todo [add|list|done] [task]")
		return
	}

	command := os.Args[1]

	switch command{
		case "add" :
			if len(os.Args)<3{
				fmt.Println("Usage: todo add [task]")
				return
			}
			task :=os.Args[2];

			err := AddTodo(task)

			if err != nil{
				fmt.Println("Error while adding the task");
			}

		case "list" : 
			err := ListTodos()
			if err != nil{
				fmt.Println("Error while adding the task");
			}

		case "done":
			if len(os.Args) < 3{
				fmt.Println("Usage: todo done [task number]")
				return
			}
		
			num, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Invalid task number")
				return
			}
			err = MarkDone(num)
			if err != nil {
				fmt.Println("Error marking task done:", err)
			}
			
		default : 
			fmt.Println("Unknown Command",command)
	}
}