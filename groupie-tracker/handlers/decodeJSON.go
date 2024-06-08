package groupieTracker

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	Id        int
	Locations []string
	Dates     string
}

type Date struct {
	Id    int
	Dates []string
}

type Relation struct {
	Id             int
	DatesLocations map[string][]string
}

func getJsonData(URL string) ([]byte, error) {
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close() // schedule Close method to be called when main function completes

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request returned status code: %d", response.StatusCode)
	}

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return dataBytes, nil
}

func getArtistsData(url string) {
	fmt.Println("Fetching Artists data...")
}
