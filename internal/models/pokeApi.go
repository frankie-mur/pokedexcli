package models

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type PokeDexLocation struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetTop20(url *string) (*PokeDexLocation, error) {
	if url == nil {
		return nil, errors.New("No more results")
	}
	resp, err := http.Get(*url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	data := PokeDexLocation{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	//set the next and prev urls to config struct

	return &data, nil
}
