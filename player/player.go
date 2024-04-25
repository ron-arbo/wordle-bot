package player

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"wordle/board"
)

// All functions from players perspective

// Should have
// find best word
// submit guess to board
// make sure to keep track of avail words

// Sample available words
var AvailableWords []string

func InitAvailableWords() {
	 // Open the file
	 // TODO: This file is gitignored
	 file, err := os.Open(wordPath)
	 if err != nil {
		 fmt.Println("Error:", err)
		 return
	 }
	 defer file.Close()
 
	 // Create a scanner to read the file line by line
	 scanner := bufio.NewScanner(file)
 
	 // Scan through the file line by line
	 for scanner.Scan() {
		 // Split the line into words
		 lineWords := strings.Fields(scanner.Text())
 
		 // Append each word to the words slice
		AvailableWords = append(AvailableWords, lineWords...)
	 }
 
	 // Check for any errors during scanning
	 if err := scanner.Err(); err != nil {
		 fmt.Println("Error:", err)
		 return
	 }
}


func FindBestWord(wordsAvailable []string) string {
	var bestWord string
	var bestScore float32

	for _, word := range wordsAvailable {
		curScore, err := wordScore(word)
		if err != nil {
			// TODO: Rather not have an error
			fmt.Println("asdf")
		}

		if curScore > bestScore {
			bestScore = curScore
			bestWord = word
		}
	}

	return bestWord
}

// TODO: Improve. Just adding up letter scores..
func wordScore(word string) (float32, error) {
	if len(word) != 5 {
		return 0, fmt.Errorf("invalid length")
	}
	word = strings.ToUpper(word)

	var wordScore float32
	seen := make(map[rune]int)
	for _, char := range word {
		times, exists := seen[char]
		if !exists {
			wordScore += letterScores[char]
			seen[char] = 1
		} else {
			// If we see a duplicate only add a fraction of the letter score
			wordScore += (letterScores[char]/(2*float32(times)))
			seen[char] += 1
		}
	}

	return wordScore, nil
}

// Assume hard mode on real wordle...
func isValidPlay(b board.Board, word string) bool {
	if len(word) != 5 {
		return false
	}
	word = strings.ToUpper(word)
	guess := []rune(strings.ToUpper(word))


	// TODO: I'm sure there's some errors with yellow....
	for i, spot := range b.Spots {
		switch spot.Color {
		case board.Green:
			// Word must have same letter in same spot
			if guess[i] != spot.Letter {
				// fmt.Printf("Invalid because %s at index %d is not %s\n", word, i, string(spot.Letter))
				return false
			}
		case board.White:
			// Word cannot contain letter
			if strings.ContainsRune(word, spot.Letter) {
				// fmt.Printf("Invalid because %s contains %s\n", word, string(spot.Letter))
				return false
			}
		case board.Yellow:
			// Word must contain letter, but not in same spot
			if !strings.ContainsRune(word, spot.Letter) || guess[i] == spot.Letter {
				// fmt.Printf("Invalid because %s does not contain %s\n", word, string(spot.Letter))
				return false
			}
		}
	}

	return true
}

func FilterAvailableWords(b board.Board) {
	var tmp []string
	for _, word := range AvailableWords {
		if isValidPlay(b, word) {
			tmp = append(tmp, word)
		}
	}
	AvailableWords = tmp
}