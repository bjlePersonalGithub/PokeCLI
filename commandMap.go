package main

import (
	"fmt"

	"github.com/bjlePersonalGithub/PokeCLI/internal/pokeapi"
)

func commandMap() error {

	err := pokeapi.LoadMaps(&displayData)
	if err != nil {
		return err
	}

	for _, area := range displayData.CurrentMapData {
		fmt.Println(area)
	}

	return nil
}
