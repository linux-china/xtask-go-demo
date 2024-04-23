package main

import (
	"fmt"
	"log"
	"os"
)

func Setup() error {
	fmt.Println("setup")
	return nil
	//return exec.Command("go", "install", "<some_tool>").Run()
}

func Hello() error {
	fmt.Println("Hello world!")
	return nil
}

func main() {
	log.SetFlags(0)
	if len(os.Args) == 1 { // print help and tasks
		log.Println("Tasks:")
		log.Println("    setup - environment setup")
	} else {
		taskName := os.Args[1]
		task, ok := map[string]func() error{
			"setup": Setup,
			"hello": Hello,
		}[taskName]
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
