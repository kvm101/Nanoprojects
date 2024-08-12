package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Path> ")
		scanner.Scan()

		path := scanner.Text()
		if path == "exit" {
			break
		}
		data, err := os.ReadFile(path)

		if err != nil {
			fmt.Println("not correct input")
			continue
		}

		commands := []string{}

		end := 0
		for i, value := range data {
			if value == 10 {
				commands = append(commands, string(data[end:i]))
				end = i + 1
			}
			if len(data) == i+1 {
				commands = append(commands, string(data[end:i]))
			}
		}

		cmd := []*exec.Cmd{}
		for _, value := range commands {
			parts := strings.Fields(value)
			cmd = append(cmd, exec.Command(parts[0], parts[1:]...))
		}

		for i, value := range cmd {
			err := value.Run()
			if err != nil {
				fmt.Printf("Command %d. %s\n", i+1, err)
			} else {
				fmt.Printf("Command %d. Success!\n", i+1)
			}
		}
	}
}
