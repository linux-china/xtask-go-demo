package main

import (
	"fmt"
	"log"
	"os"
)

func Setup() error {
	fmt.Println("setup")
	//return exec.Command("go", "install", "<some_tool>").Run()
	return nil
}

func Hello() error {
	fmt.Println("Hello world!")
	return nil
}

func Help() error {
	fmt.Println("Tasks:")
	fmt.Println("    setup - environment setup")
	fmt.Println("    hello - hello task")
	return nil
}

func AllTasks() map[string]func() error {
	return map[string]func() error{
		"setup":  Setup,
		"hello":  Hello,
		"--help": Help,
	}
}

func main() {
	log.SetFlags(0)
	if len(os.Args) == 1 { // print help and tasks
		_ = Help()
	} else {
		taskName := os.Args[1]
		allTask := AllTasks()
		task, ok := allTask[taskName]
		if !ok { // task not found
			log.Fatalln("task not found: ", taskName)
		}
		// run task
		err := task()
		if err != nil {
			log.Fatalln(err)
		}
	}

}
