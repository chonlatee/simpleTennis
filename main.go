package main

import (
	"fmt"

	"github.com/chonlatee/simpleTennis/tennis"
)

type scoreSuit struct {
	scoreA int
	scoreB int
}

func main() {

	scores := []scoreSuit{
		// A WIN
		{
			scoreA: 0,
			scoreB: 0,
		},
		{
			scoreA: 15,
			scoreB: 0,
		},
		{
			scoreA: 30,
			scoreB: 0,
		},
		{
			scoreA: 40,
			scoreB: 0,
		},
		// B WIN
		{
			scoreA: 0,
			scoreB: 15,
		},
		{
			scoreA: 15,
			scoreB: 15,
		},
		{
			scoreA: 15,
			scoreB: 30,
		},
		{
			scoreA: 15,
			scoreB: 40,
		},
	}

	for _, v := range scores {
		fmt.Printf("Score Player A: %d\t-\tPlayer B: %d\t", v.scoreA, v.scoreB)
		res, _ := tennis.ShowScore(v.scoreA, v.scoreB)
		fmt.Println(res)
	}
}
