package client

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	toolsUrl = "https://climan.dev/tools.json"
)

func GetTools() (tools []Tool, err error) {
	response, err := http.Get(toolsUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&tools)
		if err != nil {
			return nil, err
		}
		return tools, nil
	}
	return nil, errors.New("error occured while fetching tools.json")
}
