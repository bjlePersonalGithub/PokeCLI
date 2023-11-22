package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type mapData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func main() {
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		userInput := scanner.Text()
		commands := getCLICommands()

		if command, commandExist := commands[userInput]; commandExist {
			commandName := strings.ToLower(command.name)
			if commandName == "help" {
				fmt.Println("\nWelcome to the Pokedex!\n" + "Usage:\n")
				for _, com := range commands {
					fmt.Println(com.name + ": " + com.description)
				}
			}
			if commandName == "exit" {
				return
			}
			if commandName == "map" {
				commandMap()
			}
			fmt.Println()
		} else {
			fmt.Println("Invalid Command")
		}

	}
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
		}, "map": {
			name:        "map",
			description: "Display the map",
			callback:    commandMap,
		},
	}
}

func commandHelp() error {
	// Your implementation here
	return nil
}

func commandExit() error {
	return errors.New("Exit Program")
}

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}
	mapJson := mapData{}
	umarshalErr := json.Unmarshal(body, &mapJson)

	if umarshalErr != nil {
		return umarshalErr
	}

	fmt.Println(mapJson.Results)
	return nil
}
