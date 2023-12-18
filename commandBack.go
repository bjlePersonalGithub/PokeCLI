package main

import "fmt"

func commandMapBack() error {
	if displayData.PreviousMapData == nil || len(displayData.PreviousMapData) == 0 {
		return fmt.Errorf("no previous map data")
	}

	return nil
}
