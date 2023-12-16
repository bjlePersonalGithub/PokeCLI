package pokeapi

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetInitialDisplayMap() mapDisplayData {
	return mapDisplayData{
		currentMapData:  make([]string, 0, 20),
		previousMapData: make([]string, 0, 20),
		locationIndex:   1,
	}
}

func LoadMaps(mapData mapDisplayData) ([]string, error) {
	res, err := http.Get(locationURL + strconv.Itoa(mapData.locationIndex))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	fmt.Println(res)
	return nil, nil
}
