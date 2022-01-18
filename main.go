package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type Artist struct{
	Name string json:"id"
	Image []byte
	Members string
	DoA string
	DoFA string
}

type Location struct{

}

func homePage (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Homepage")
}

func main() {
	fmt.Println("Starting Server at Port 8080")
	fmt.Println("now open a broswer and enter: localhost:8080 into the URL")
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
