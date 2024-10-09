package main

import (
	"fmt"
	"log"
)

func main() {

	testPlayers := []*Player{
		{
			Name: "Leandro",
		},
		{
			Name: "Antonio",
		},
	}

	board, err := NewBoard(testPlayers)
	if err != nil {
		log.Fatalf("could not init board: %v", err)
	}
	for _, v := range board.GetDeck().Shuffle().Cards {
		fmt.Println(v)
	}
}
