package main

import (
	"github.com/bjlePersonalGithub/PokeCLI/internal/pokeapi"
)

func commandMap() error {

	displayData := pokeapi.GetInitialDisplayMap()

	pokeapi.LoadMaps(displayData)

	// if mapState.nextURL == "" {
	// 	mapState.nextURL = "https://pokeapi.co/api/v2/location/"
	// }

	// if mapState.nextURL != "" {
	// 	mapState.prevURL = mapState.nextURL
	// }

	// res, err := http.Get(mapState.nextURL)
	// if err != nil {
	// 	return err
	// }
	// defer res.Body.Close()

	// if res.StatusCode != http.StatusOK {
	// 	return fmt.Errorf("request failed with status code: %d", res.StatusCode)
	// }

	// var mapData mapData
	// if err := json.NewDecoder(res.Body).Decode(&mapData); err != nil {
	// 	return err
	// }

	// for _, result := range mapData.Results {
	// 	fmt.Println(result.Name)
	// }

	// mapState.nextURL = mapData.Next // Update the nextURL for the next fetch
	// mapState.prevURL = mapData.Previous

	return nil
}
