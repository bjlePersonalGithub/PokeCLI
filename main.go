package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
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
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var mapState = struct {
	nextURL string
	prevURL string
}{}

func main() {
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
	return scanner.Text()
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

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!\n" + "Usage:\n")
	for _, com := range getCLICommands() {
		fmt.Println(com.name + ": " + com.description)
	}
	return nil
}
func commandExit() error {
	return errors.New("exit program")
}

func commandMap() error {
	if mapState.nextURL == "" {
		mapState.nextURL = "https://pokeapi.co/api/v2/location/"
	}

	if mapState.nextURL != "" {
		mapState.prevURL = mapState.nextURL
	}

	res, err := http.Get(mapState.nextURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	var mapData mapData
	if err := json.NewDecoder(res.Body).Decode(&mapData); err != nil {
		return err
	}

	for _, result := range mapData.Results {
		fmt.Println(result.Name)
	}

	mapState.nextURL = mapData.Next // Update the nextURL for the next fetch

	return nil
}

func commandMapBack() error {
	if mapState.prevURL == "" {
		fmt.Println("You're already at the first page")
		return nil
	}

	res, err := http.Get(mapState.prevURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	var mapData mapData
	if err := json.NewDecoder(res.Body).Decode(&mapData); err != nil {
		return err
	}

	for _, result := range mapData.Results {
		fmt.Println(result.Name)
	}

	mapState.nextURL = mapState.prevURL // Store current URL as next
	mapState.prevURL = mapData.Previous // Update the prevURL for the previous fetch

	return nil
}
