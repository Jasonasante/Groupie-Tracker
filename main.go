package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Relation struct {
	Artist
}

// type Artist struct {
// 	Id              uint   `json:"id"`
// 	Name            string `json:"name"`
// 	Image           string `json:"image"`
// 	Members         string `json:"members"`
// 	CreationDate    uint   `json:"creationDate"`
// 	FirstAlbum      string `json:"firstAlbum"`
// 	ConcertLocations string `json:"locations"`
// 	ConcertDates    string `json:"concertDates"`
// 	Relations string
// }

type Artist struct {
	Id               uint
	Name             string
	Image            string
	Members          string
	CreationDate     uint
	FirstAlbum       string
	ConcertLocations string
	ConcertDates     string
	Relations        string
}

var Relationship []Relation

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Homepage")

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(Relationship)
// }

func main() {
	artist, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	artistData, _ := ioutil.ReadAll(artist.Body)
	// fmt.Println(string(artistData))
	// We made type Artist a slice because our api is an array of json data
	var artistInfo []Artist
	json.Unmarshal(artistData, &artistInfo)
	for _, name := range artistInfo {
		fmt.Println(name.Name)
	}
	// fmt.Println((jartist))

	// fmt.Println("Starting Server at Port 8080")
	// fmt.Println("now open a broswer and enter: localhost:8080 into the URL")
	// http.HandleFunc("/", homePage)
	// http.ListenAndServe(":8080", nil)
}
