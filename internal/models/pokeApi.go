package models

import (
	"encoding/json"
	"fmt"
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

type PokeDexLocationAreas struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"encounter_method,omitempty"`
		VersionDetails []struct {
			Rate    int `json:"rate,omitempty"`
			Version struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version,omitempty"`
		} `json:"version_details,omitempty"`
	} `json:"encounter_method_rates,omitempty"`
	GameIndex int `json:"game_index,omitempty"`
	ID        int `json:"id,omitempty"`
	Location  struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"location,omitempty"`
	Name  string `json:"name,omitempty"`
	Names []struct {
		Language struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"language,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"names,omitempty"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"pokemon,omitempty"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance,omitempty"`
				ConditionValues []any `json:"condition_values,omitempty"`
				MaxLevel        int   `json:"max_level,omitempty"`
				Method          struct {
					Name string `json:"name,omitempty"`
					URL  string `json:"url,omitempty"`
				} `json:"method,omitempty"`
				MinLevel int `json:"min_level,omitempty"`
			} `json:"encounter_details,omitempty"`
			MaxChance int `json:"max_chance,omitempty"`
			Version   struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version,omitempty"`
		} `json:"version_details,omitempty"`
	} `json:"pokemon_encounters,omitempty"`
}

func GetTop20(c *Cache, url *string) (*PokeDexLocation, error) {
	if url == nil {
		return nil, errors.New("No more results")
	}
	//Check if the url is in the cache if so return
	val, exists := c.Get(*url)
	if exists {
		data, err := toPokeDexJson(val)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	fmt.Println("Making request...")

	resp, err := http.Get(*url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)

	fmt.Println("Received response")

	resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("About to add to cache")
	//Add the response to the cache with key being the url
	c.Add(*url, body)

	fmt.Printf("Finished adding")

	data, err := toPokeDexJson(body)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetNamesFromArea(c *Cache, name string) (*PokeDexLocationAreas, error) {
	if name == "" {
		return nil, errors.New("Please provide a name")
	}
	//Check cache
	val, exists := c.Get(name)
	if exists {
		data, err := toPokeDexLocationAreaJson(val)
		if err != nil {
			return nil, err
		}
		return data, nil

	}

	const url = "https://pokeapi.co/api/v2/location-area"
	urlWithName := fmt.Sprintf("%s/%s", url, name)

	resp, err := http.Get(urlWithName)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)

	fmt.Println("Received response")

	resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	//Add the response to the cache with key being the name
	c.Add(name, body)

	fmt.Printf("Finished adding")

	data, err := toPokeDexLocationAreaJson(body)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func toPokeDexJson(body []byte) (*PokeDexLocation, error) {
	data := PokeDexLocation{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func toPokeDexLocationAreaJson(body []byte) (*PokeDexLocationAreas, error) {
	data := PokeDexLocationAreas{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
