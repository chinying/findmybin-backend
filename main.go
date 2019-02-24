package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/levigross/grequests"
)

type MapboxRequest struct {
	Expires string   `json:"expires`
	Scopes  []string `json:"scopes"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MAPBOX_USERNAME := os.Getenv("MAPBOX_USERNAME")
	MAPBOX_ACCESS_TOKEN := os.Getenv("MAPBOX_ACCESS_TOKEN")
	url := fmt.Sprintf("https://api.mapbox.com/tokens/v2/%s?access_token=%s", MAPBOX_USERNAME, MAPBOX_ACCESS_TOKEN)
	fmt.Println(url)

	resp, err := grequests.Post(url, nil)
	// resp, err := grequests.Get("https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	fmt.Println(resp.String())
}
