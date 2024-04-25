package main

import (
	"fmt"
	"wordle/board"
	"wordle/player"

	"github.com/fatih/color"
)

const (
	green  = color.FgGreen
	white  = color.FgWhite
	yellow = color.FgYellow
)


// Sample board
var board1 = board.Board{
	Spots: []*board.Spot{
		{Letter: 'P', Color: yellow},
		{Letter: 'L', Color: yellow},
		{Letter: 'A', Color: yellow},
		{Letter: 'N', Color: yellow},
		{Letter: 'E', Color: white},
	},
}

const (
	// TODO: Make this a CLI flag
	interactive = true
)

func main() {

	// Loop structure
	// 1. Player guess
	// 2. Board fields guess, returns new board
	// 3. Player sees board, creates new guess
	// 4. Repeat until board recognizes win (or player)?

	fmt.Println("Answer is ", board.Answer)

	// Blank board
	b := board.Board{
		Spots: []*board.Spot{
		{Letter: 0, Color: board.White},
		{Letter: 0, Color: board.White},
		{Letter: 0, Color: board.White},
		{Letter: 0, Color: board.White},
		{Letter: 0, Color: board.White},
		},
	}
	var guess string

	player.InitAvailableWords()

	// TODO: Limit to certain # of guesses
	for !b.IsSolved() {
		if interactive {
			fmt.Print("Guess: ")
			fmt.Scanln(&guess)
			// TODO: Validate guess
		} else {
			player.FilterAvailableWords(b)
			guess = player.FindBestWord(player.AvailableWords) 
		}

		fmt.Println("Player guesses ", guess)
		b = *board.BuildBoard(guess)
		b.Print()
	}
}
