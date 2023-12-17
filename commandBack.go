package main

func commandMapBack() error {
	// if mapState.prevURL == "" {
	// 	fmt.Println("You're already at the first page")
	// 	return nil
	// }

	// res, err := http.Get(mapState.prevURL)
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

	// mapState.nextURL = mapState.prevURL // Store current URL as next
	// mapState.prevURL = mapData.Previous // Update the prevURL for the previous fetch

	return nil
}
