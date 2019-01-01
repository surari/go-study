package main

import "fmt"

type sushi struct {
	name  string
	score int
}

type user struct {
	name string
	data []sushi
}

func main() {
	meryu := user{
		name: "meryu",
		data: []sushi{
			{name: "tuna", score: 5},
			{name: "ikura", score: 1},
			{name: "tai", score: 2},
		},
	}
	sara := user{
		name: "sara",
		data: []sushi{
			{name: "tuna", score: 2},
			{name: "ikura", score: 2},
			{name: "tai", score: 2},
		},
	}
	karusya := user{
		name: "karusya",
		data: []sushi{
			{name: "tuna", score: 2},
			{name: "ikura", score: 10},
			{name: "tai", score: 5},
		},
	}
	mari := user{
		name: "karusya",
		data: []sushi{
			{name: "tuna", score: 2},
			{name: "ikura", score: 6},
			{name: "tai", score: 0},
		},
	}
	fmt.Println(meryu)
	fmt.Println(sara)
	fmt.Println(karusya)
	fmt.Println(mari)
}
