package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/levigross/grequests"
)

type MapboxRequest struct {
	Expires string   `json:"expires"`
	Scopes  []string `json:"scopes"`
}

type FakeRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
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

	// to format to 2018-09-20T06:30:00Z
	expiry := time.Now().UTC().Add(time.Minute * 30)
	fmt.Println(expiry.Format("2006-01-02T15:04:05Z"))
	scopes := []string{"styles:tiles", "styles:read", "fonts:read", "datasets:read"}
	requestJSON := MapboxRequest{expiry.Format("2006-01-02T15:04:05Z"), scopes}

	resp, err := grequests.Post(url, &grequests.RequestOptions{JSON: requestJSON})
	// resp, err := grequests.Get("https://jsonplaceholder.typicode.com/todos/1", &grequests.RequestOptions{JSON: MapboxRequest{}})
	// resp, err := grequests.Post("https://jsonplaceholder.typicode.com/posts", &grequests.RequestOptions{JSON: FakeRequest{"hello", "body", 4}, IsAjax: true})
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	fmt.Println(resp.String())
}
