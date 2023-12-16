package pokeapi

func GetDisplayMap() mapDisplayData {
	return mapDisplayData{
		currentMapData:  make([]string, 0, 20),
		previousMapData: make([]string, 0, 20),
	}
}

func LoadInitialMaps(data mapDisplayData) []string {

	return nil
}
