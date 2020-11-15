package  main

import (
	"fmt"
	"net/http"
	"encoding/json"
)


func main() {

	fmt.Println("server is running at port 3333")
	http.HandleFunc("/", homepage)
	http.HandleFunc("/about", about)
	http.HandleFunc("/api/music", music)
	http.HandleFunc("/api/album", album)
	
	http.ListenAndServe(":3333", nil)
}

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "this is home page")
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "this is about page")
}

func music(w http.ResponseWriter, r *http.Request){
	data := Music{
		1,
		"only love",
		"trade mark",
	}

	json.NewEncoder(w).Encode(data)
}

func album(w http.ResponseWriter, r *http.Request){
	listmusic := Album{
		Music{1,
		"only love",
		"trade mark",
		},
		Music{
			2,
			"missing you",
			"lyly",
		},
		Music{
			2,
			"bigcityboy",
			"MTP",
		},

	}

	json.NewEncoder(w).Encode(listmusic)
}


type Music struct {

	Id 		 int		`json:"id"`
	Name	 string		`json:"name"`
	Singer 	 string 	`json:"singer"`
}


type Album []Music