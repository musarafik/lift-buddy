package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"time"
)

type Exercise struct {
	Name   string
	Reps   int
	Sets   int
	Weight int
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	currDayOfWeek := time.Now().Weekday()
	exercises := []Exercise{}

	for range rand.IntN(10) {
		exercises = append(exercises, getRandomExercise())
		fmt.Println(exercises)
	}

	rootComponent := root(currDayOfWeek.String(), exercises)

	rootComponent.Render(r.Context(), w)
}

func main() {
	http.HandleFunc("/", rootHandler)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
