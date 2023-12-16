package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type mapData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var mapState = struct {
	nextURL string
	prevURL string
}{}
