package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, com := range getCLICommands() {
		fmt.Println(com.name + ": " + com.description)
	}
	return nil
}
