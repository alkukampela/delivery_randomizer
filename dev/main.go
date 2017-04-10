package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Randomizable struct {
	Key		string
	Value		float64
	Min 		float64
	Max 		float64
	Variance 	float64
}

type Randomized struct {
	Key		string	`json:"key"`
	Value		float64	`json:"value"`
}

func handler(writer http.ResponseWriter, req *http.Request) {
	var randomizables []Randomizable
	var results []Randomized

	if req.Body == nil {
		http.Error(writer, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(req.Body).Decode(&randomizables)
	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	for _, randomizable := range randomizables {
		results = append(results, randomize(randomizable))
	}

	resultJson, err := json.Marshal(results)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	fmt.Fprintf(writer, string(resultJson))
}

func randomize(randomizable Randomizable) Randomized {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := r.NormFloat64() * randomizable.Variance + randomizable.Value

	if result > randomizable.Max {
		result = randomizable.Max
	}

	if result < randomizable.Min {
		result = randomizable.Min
	}

	return Randomized { randomizable.Key, result }
}

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable was not set")
	}

	log.Println("Attempting to listen on port: ", port)
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		log.Fatal("Could not listen: ", err)
	}
}
