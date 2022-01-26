package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
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

// 	DataMap := make([][]string, len(relationInfo))
// 	for i := range DataMap {
// 		DataMap[i] = make([]string, len(relationInfo))
// 	}
// count:=0
// 	for j, ele := range relationInfo {
// 		for i, mapData := range ele.DatesLocations {
// 			DataMap[j]=ele.DatesLocations[i]
// 			DataMap[j][count]=mapData[count]
// 			count++
// 		}
// 	}
// 	fmt.Println(DataMap)
	return relationInfo
}

func collectData() []Data {
	ArtistData()
	RelationData()
	dataData := make([]Data, len(artistInfo))
	for i := 0; i < len(artistInfo); i++ {
		dataData[i].A = artistInfo[i]
		dataData[i].R = relationInfo[i]
	}
	return dataData
}

func HandleRequests() {
	fmt.Println("Starting Server at Port 8080")
	fmt.Println("now open a broswer and enter: localhost:8080 into the URL")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/artistInfo", artistPage)
	http.HandleFunc("/locations", returnAllLocations)
	http.HandleFunc("/dates", returnAllDates)
	http.HandleFunc("/relation", returnAllRelation)
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArtists")
	data := ArtistData()
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, data)
}

func artistPage(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("ArtistName")
	fmt.Println(value)
	a := collectData()
	var b Data
	for i, ele := range collectData() {
		if value == ele.A.Name {
			b = a[i]
		}
	}
	t, _ := template.ParseFiles("artistPage.html")
	t.Execute(w, b)
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

func main() {
	HandleRequests()
}
