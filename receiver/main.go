package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var rd []*ReceiveData

type ReceiveData struct {
	Value       string    `json:"value"`
	LastUpdated time.Time `json:"lastUpdated"`
}

func main() {
	go Monitor()
	r := mux.NewRouter()
	r.HandleFunc("/rate", Rate).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func Rate(w http.ResponseWriter, r *http.Request) {
	receiveData := ReceiveData{}
	err := json.NewDecoder(r.Body).Decode(&receiveData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	receiveData.LastUpdated = time.Now()
	rd = append(rd, &receiveData)
	return
}

func Monitor() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		if rd != nil {
			if len(rd) == 1 {
				if time.Since(rd[0].LastUpdated) > 30*time.Second {
					fmt.Println("Warning")
				}
			}
			if len(rd) > 1 {
				if rd[len(rd)-1].LastUpdated.Sub(rd[len(rd)-2].LastUpdated) > 30*time.Second {
					fmt.Println("Warning")
				} else {
					fmt.Println("Good")
				}
			}
		}
		<-ticker.C
	}
}
