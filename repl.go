package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func promptREPL() {
	for {
		fmt.Print("Pokedex > ")
		userInput := getUserInput()
		commands := getCLICommands()
		command := commands[userInput]
		switch strings.ToLower(command.name) {
		case "help":
			if err := commandHelp(); err != nil {
				fmt.Println("Error displaying help:", err)
			}
		case "map":
			if err := commandMap(); err != nil {
				fmt.Println("Error fetching maps:", err)
			}
		case "mapb":
			if err := commandMapBack(); err != nil {
				fmt.Println("Error fetching previous maps:", err)
			}
		case "exit":
			return
		default:
			fmt.Println("Invalid Command")
		}

	}
}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return ""
	}
	return strings.ToLower(scanner.Text())
}

func getCLICommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 maps",
			callback:    commandMapBack,
		},
	}
}
