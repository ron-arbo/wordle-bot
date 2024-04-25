package board

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// Contains all functions needed from boards perspective

// Should have
// validate user play
// check guess
// return new board
// see if won, return stats
// keep track of board state


const (
	Green  = color.FgGreen
	White  = color.FgWhite
	Yellow = color.FgYellow
)

type Board struct {
	Spots []*Spot
}

type Spot struct {
	Letter rune
	Color  color.Attribute
}

var (
	Answer = "TRADE"
	answerMap = make(map[rune]int)
	runeAnswer = []rune(Answer)
)

func initAnswerMap() {
	for _, letter := range runeAnswer {
		_, exists := answerMap[letter]
		if !exists {
			answerMap[letter] = 1
		} else {
			answerMap[letter] += 1
		}
	}
}

// Print will print the current board (Assumes all caps)
func (b Board) Print() {
	if len(b.Spots) != 5 {
		// TODO: err?
		fmt.Println("BAD BOARD")
		return
	}

	for i := 0; i < 7; i++ {
		for _, spot := range b.Spots {
			bigLetter, exists := bigLetters[spot.Letter]
			if !exists {
				// TODO: err?
				fmt.Println("BAD CHAR: ", bigLetter)
				return
			}

			colorPrint(bigLetter[i]+"  ",spot.Color)
		}
		fmt.Print("\n")
	}

}

// BuildBoard will create a board from the given guess
func BuildBoard(guess string) *Board {
	// TODO: Validate guess, answer?
	guess = strings.ToUpper(guess)
	runeGuess := []rune(guess)

	// Initialize with the letters
	// TODO: Do this in below loop?
	spots := []*Spot{
		{Letter: runeGuess[0]},
		{Letter: runeGuess[1]},
		{Letter: runeGuess[2]},
		{Letter: runeGuess[3]},
		{Letter: runeGuess[4]},
	}

	initAnswerMap()

	// Update colors
	for i, spot := range spots {
		// Green
		if spot.Letter == runeAnswer[i] {
			spot.Color = Green
			answerMap[spot.Letter] -= 1
		}

		// White
		_, exists := answerMap[spot.Letter]
		if !exists {
			spot.Color = White
			continue
		}
	}	

	// Update yellows last. Can't update answerMap too early for a misplaced yellow. 
	// If green letter is at index 5 but we see the same letter at index 0, we would incorrectly mark that yellow and decrement answerMap
	for _, spot := range spots {
		// Ignore already decided colors
		// TODO: What is 0 value for colors?
		if spot.Color == Green || spot.Color == White {
			continue
		}

		// Yellow
		remaining := answerMap[spot.Letter]
		if remaining > 0 {
			spot.Color = Yellow
			// Update answerMap to keep track of dupes
			answerMap[spot.Letter] -= 1
		}	
	}	

	return &Board{
		Spots: spots,
	}
}

// IsSolved will return a boolean indicating if the given board is solved
func (b Board) IsSolved() bool {
	for _, spot := range b.Spots {
		if spot.Color != Green {
			return false
		}
	}
	return true
}


func colorPrint(str string, color color.Attribute) {
	fmt.Printf("\x1b[%dm%s\x1b[0m", color, str)
}
