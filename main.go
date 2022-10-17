package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	element int    `json:"water"`
	status  string `json:"status"`
}

var PORT = ":9000"

func main() {
	http.HandleFunc("/", ElementStatus)

	http.ListenAndServe(PORT, nil)

}

func ElementStatus(w http.ResponseWriter, r *http.Request) {
	ticker := time.NewTicker(15 * time.Second)
	for _ = range ticker.C {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {

			rand.Seed(time.Now().UnixNano())
			min := 1
			max := 100
			water := rand.Intn(max-min+1) + min
			wind := rand.Intn(max-min+1) + min

			if water < 5 {
				data := Status{status: "Status Water Aman", element: water}
				fmt.Println(data)
			} else if water >= 6 && water <= 8 {
				data := Status{status: "Status Water Siaga", element: water}
				fmt.Println(data)
			} else {
				data := Status{status: "Status Water Bahaya", element: water}
				fmt.Println(data)
			}

			if wind < 6 {
				data := Status{status: "Status Wind Aman", element: wind}
				fmt.Println(data)
			} else if wind >= 7 && wind <= 15 {
				data := Status{status: "Status Wind Siaga", element: wind}
				fmt.Println(data)
			} else {
				data := Status{status: "Status Wind Bahaya", element: wind}
				fmt.Println(data)
			}
			// json.NewDecoder(w).Encode(Status)
			return
		}
		http.Error(w, "Invalid", http.StatusBadRequest)
	}

}
