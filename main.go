package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"text/template"
)

type Data struct {
	A Artist
	R Relation
}

type Artist struct {
	Id           uint     `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate uint     `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Location struct {
	Id        uint     `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	Id    uint     `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type GetImage struct {
	Img []string
}

var (
	artistInfo   []Artist
	locationMap  map[string]json.RawMessage
	locationInfo []Location
	datesMap     map[string]json.RawMessage
	datesInfo    []Date
	relationMap  map[string]json.RawMessage
	relationInfo []Relation
)

func ArtistData() []Artist {
	artist, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	artistData, _ := ioutil.ReadAll(artist.Body)
	json.Unmarshal(artistData, &artistInfo)
	return artistInfo
}

func LocationData() []Location {
	var bytes []byte
	location, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	locationData, _ := ioutil.ReadAll(location.Body)
	err := json.Unmarshal(locationData, &locationMap)
	if err != nil {
		fmt.Println("error :", err)
	}
	for _, m := range locationMap {
		for _, v := range m {
			bytes = append(bytes, v)
		}
	}
	err = json.Unmarshal(bytes, &locationInfo)
	if err != nil {
		fmt.Println("error :", err)
	}
	fmt.Println(locationInfo[1])
	return locationInfo
}

func DatesData() []Date {
	var bytes []byte
	dates, _ := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	datesData, _ := ioutil.ReadAll(dates.Body)
	err := json.Unmarshal(datesData, &datesMap)
	if err != nil {
		fmt.Println("error :", err)
	}
	for _, m := range datesMap {
		for _, v := range m {
			bytes = append(bytes, v)
		}
	}
	err = json.Unmarshal(bytes, &datesInfo)
	if err != nil {
		fmt.Println("error :", err)
	}
	return datesInfo
}

func RelationData() []Relation {
	var bytes []byte
	relation, _ := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	relationData, _ := ioutil.ReadAll(relation.Body)
	err := json.Unmarshal(relationData, &relationMap)
	if err != nil {
		fmt.Println("error :", err)
	}

	for _, m := range relationMap {
		for _, v := range m {
			bytes = append(bytes, v)
		}
	}
	err = json.Unmarshal(bytes, &relationInfo)
	if err != nil {
		fmt.Println("error :", err)
	}
	return relationInfo
}

func collectData() {
	dataData := make([]Data, len(artistInfo))

	for i := 0; i < len(artistInfo); i++ {
		for _, ele:=range dataData{
			ele.A=append(ele[i].A, artistInfo[i])
		}
		
		//dataData[i].R = append(dataData[i].R, relationInfo[i])
	}
	fmt.Println(dataData)
}

func HandleRequests() {
	fmt.Println("Starting Server at Port 8080")
	fmt.Println("now open a broswer and enter: localhost:8080 into the URL")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/artists", returnAllArtists)
	http.HandleFunc("/locations", returnAllLocations)
	http.HandleFunc("/dates", returnAllDates)
	http.HandleFunc("/relation", returnAllRelation)
	http.ListenAndServe(":8080", nil)
}

func returnAllArtists(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArtists")
	json.NewEncoder(w).Encode(ArtistData())
}

func returnAllLocations(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllLocations")
	json.NewEncoder(w).Encode(LocationData())
}

func returnAllDates(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllDates")
	json.NewEncoder(w).Encode(DatesData())
}

func returnAllRelation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllRelation")
	json.NewEncoder(w).Encode(RelationData())
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Groupie-Tracker")
}

func main() {
	collectData()
	HandleRequests()
}
