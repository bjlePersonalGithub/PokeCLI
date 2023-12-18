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
		locationIndex:   1,
		areaIndex:       0,
	}
}

func getLocationInfo(locationIndex int) (*locationData, error) {
	res, err := http.Get(locationURL + strconv.Itoa(locationIndex))
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
	resData, err := getLocationInfo(displayData.locationIndex)
	if err != nil {
		return err
	}

	//Load new areas
	for len(displayData.CurrentMapData) < 20 {
		for i := displayData.areaIndex; i < len(resData.Areas); i++ {
			// If we reach max of currentMapData capcaity
			if len(displayData.CurrentMapData) == 20 {
				break
			} else {
				displayData.CurrentMapData = append(displayData.CurrentMapData, resData.Areas[i].Name)
				displayData.areaIndex += 1
			}
			if i+1 == len(resData.Areas) {
				// If we've reached the end of the areas, increment the location
				displayData.areaIndex = 0
				displayData.locationIndex += 1
			}
		}
		// If we've reached the end of the areas, load the next location
		resData, err = getLocationInfo(displayData.locationIndex)
		if err != nil {
			return err
		}
	}
	return nil
}
