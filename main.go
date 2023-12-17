package main

import "github.com/bjlePersonalGithub/PokeCLI/internal/pokeapi"

var displayData = pokeapi.GetInitialDisplayMap()

func main() {
	promptREPL()
}
