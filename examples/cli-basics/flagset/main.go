package main

import (
	"fmt"
	"os"
	"log"
	"flag"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no command provided")
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetCmd.String("msg", "Reminders App - Cli Basics", "the message for the greet command")
		err := greetCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Printf("Hello %s\n", *msgFlag)
	case "help":
		fmt.Println("Some help")
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
	}
}
