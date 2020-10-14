package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no command provided")
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "greet":
		msg := "REMINDERS APP - CLI BASICS"
		if len(os.Args) > 2 {
			flag := strings.Split(os.Args[2], "=")
			if len(flag) == 2 && flag[0] == "--msg" {
				msg = flag[1]
			}
		}
		fmt.Printf("Hello %s\n", msg)
	case "help":
		fmt.Println("Some help")
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
	}
}
