package main

import (
    "os"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"time"
	"math/rand"
)

type Randomizable struct{
	Key			string
    Value   	float64
	Min 		float64
	Max 		float64
	Variance 	float64
}

type Randomized struct{
	Key			string
    Value   	float64
}


func handler(writer http.ResponseWriter, req *http.Request) {
	var randomizables []Randomizable
	var result []Randomized
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
		result = append(result, randomize(randomizable))
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	fmt.Fprintf(writer, string(resultJson))
}


func randomize(randomizable Randomizable) Randomized {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Randomize input
	result := r.NormFloat64() * randomizable.Variance + randomizable.Value

	if (result > randomizable.Max) {
		result = randomizable.Max
	}

	if (result < randomizable.Min) {
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
