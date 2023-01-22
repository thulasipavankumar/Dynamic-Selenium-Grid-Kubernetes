package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const static_data = `{
	"value": {
	  "sessionId": "8c2cfa36e2b38c82621437224539e47f",
	  "capabilities": {
		"acceptInsecureCerts": false,
		"browserName": "chrome",
		"browserVersion": "104.0.5112.79",
		"platformName": "LINUX"

	  }
	}
  }`

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {

}
func makeFunCall() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(responseData))
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	for _, pokemon := range responseObject.Pokemon {
		fmt.Println(pokemon.Species.Name)
	}
}
func printDataAndType(i interface{}) {
	fmt.Printf("Type is %T and value is \n", i)
}
