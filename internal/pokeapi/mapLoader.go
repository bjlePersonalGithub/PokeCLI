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
	}
}

func getNewLocationInfo(locationIndex int) (*locationData, error) {
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
		displayData.locationIndex += 1
	}

	resData, err := getNewLocationInfo(displayData.locationIndex)
	if err != nil {
		return err
	}

	for len(displayData.CurrentMapData) < 20 {
		for i := 0; i < len(resData.Areas); i++ {
			if len(displayData.CurrentMapData) == 20 {
				break
			} else {
				displayData.CurrentMapData = append(displayData.CurrentMapData, resData.Areas[i].Name)
			}
		}
		displayData.locationIndex += 1
		resData, err = getNewLocationInfo(displayData.locationIndex)
		if err != nil {
			return err
		}
	}

	return nil
}
