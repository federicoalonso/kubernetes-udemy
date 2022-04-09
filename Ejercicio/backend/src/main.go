package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type HandsOn struct {
	Time     time.Time `json:"time"`
	Hostname string    `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	resp := HandsOn{
		Time:     time.Now(),
		Hostname: os.Getenv("HOSTNAME"),
	}

	jsonResponse, err := json.Marshal(&resp)

	if err != nil {
		w.Write([]byte("Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":9090", nil)
}

// para correrlo en un contenedor, se usa el siguiente comando:
// docker run --rm -dti -v $PWD/:/go --net host --name ejercicio-go-1 golang bash
