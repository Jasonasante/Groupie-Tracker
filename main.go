package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Relation struct {
	Artist
}

type ArtistSlice struct {
	ArtistSl []Artist
}

type Artist struct {
	Id               uint     `json:"id"`
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	Members          []string `json:"members"`
	CreationDate     uint     `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	ConcertLocations string   `json:"locations"`
	ConcertDates     string   `json:"concertDates"`
	Relations        string   `json:"relations"`
	MemNames         string
	artistInf []Artist
}

var (
	Relationship []Relation
	SliceArtist  []Artist
)

func homePage(w http.ResponseWriter, r *http.Request) {
	artist, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	artistData, _ := ioutil.ReadAll(artist.Body)
	// We made type Artist a slice because our api is an array of json data
	var artistInfo []Artist
	var artistId uint
	var artistName string
	var artistImage string
	var artistMembers []string
	var artistCreationDate uint
	var artistFirstAlbum string
	var artistConcertLocations string
	var artistConcertDates string
	var artistRelations string
	var members string
	// var ArtistS ArtistSlice

	json.Unmarshal(artistData, &artistInfo)
	for _, data := range artistInfo {
		SliceArtist = append(SliceArtist, data)
		artistId = data.Id
		artistName = data.Name
		artistImage = data.Image
		artistMembers = data.Members
		artistCreationDate = data.CreationDate
		artistFirstAlbum = data.FirstAlbum
		artistConcertLocations = data.ConcertLocations
		artistConcertDates = data.ConcertDates
		artistRelations = data.Relations
		for i, mem := range artistMembers {
			if i < len(artistMembers)-1 {
				members += mem
				members += "\n"
			} else {
				members += mem
			}
		}
		t, err := template.ParseFiles("template.html")
		if err != nil {
			log.Fatal()
		}
		t.Execute(w, Artist{Id: artistId, Name: artistName, Image: artistImage, Members: artistMembers, CreationDate: artistCreationDate, FirstAlbum: artistFirstAlbum, ConcertLocations: artistConcertLocations, ConcertDates: artistConcertDates, Relations: artistRelations, MemNames: members})
	}
	fmt.Println(artistInfo[1])
	fmt.Println(len(artistInfo))
}

func main() {
	// fmt.Println((jartist))

	fmt.Println("Starting Server at Port 8080")
	fmt.Println("now open a broswer and enter: localhost:8080 into the URL")
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
