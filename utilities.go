package main

import "math/rand/v2"

func getRandomExercise() Exercise {
	names := []string{"bench press", "shoulder press", "squat", "deadlift"}

	return Exercise{
		Name:   names[rand.IntN(len(names)-1)],
		Weight: rand.IntN(500),
		Sets:   rand.IntN(5),
		Reps:   rand.IntN(12),
	}
}
