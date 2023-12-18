package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetInitialDisplayMap() mapDisplayData {
	return mapDisplayData{
		CurrentMapData:  make([]string, 0, 20),
		PreviousMapData: nil,
		LocationIndex:   1,
		AreaIndex:       0,
	}
}

func getLocationInfo(LocationIndex int) (*locationData, error) {
	res, err := http.Get(locationURL + strconv.Itoa(LocationIndex))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resData *locationData
	if err := json.NewDecoder(res.Body).Decode(&resData); err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	return resData, nil
}

func LoadMaps(displayData *mapDisplayData) error {

	if len(displayData.CurrentMapData) == 20 {
		displayData.PreviousMapData = displayData.CurrentMapData
		displayData.CurrentMapData = make([]string, 0, 20)
	}

	// Load new location
	resData, err := getLocationInfo(displayData.LocationIndex)
	if err != nil {
		return err
	}

	//Load new areas
	for len(displayData.CurrentMapData) < 20 {
		for i := displayData.AreaIndex; i < len(resData.Areas); i++ {
			// If we reach max of currentMapData capcaity
			if len(displayData.CurrentMapData) == 20 {
				break
			} else {
				displayData.CurrentMapData = append(displayData.CurrentMapData, resData.Areas[i].Name)
				displayData.AreaIndex += 1
			}
			if i+1 == len(resData.Areas) {
				displayData.AreaIndex = 0
				displayData.LocationIndex += 1
			}
		}
		// If we've reached the end of the areas, load the next locatio
		resData, err = getLocationInfo(displayData.LocationIndex)
		if err != nil {
			return err
		}
	}
	return nil
}
