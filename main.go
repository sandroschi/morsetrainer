package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/oto"
)

func main() {
	// Read config
	var config = read_config()
	var morse = NewMorseCharacterPool(config)

	// Prepare sound output
	ctx, _ := oto.NewContext((int)(config.SampleRate), 1, 2, 0)
	player := ctx.NewPlayer()
	defer player.Close()

	// Prepare user input
	reader := bufio.NewReader(os.Stdin)
	var input string = "j"
	var trainingDuration time.Duration = 0

	for input == "j" { // User wants training
		// Generate random characters and convert them to morse symbols
		characters := morse.GetRandomCharacters(config.CharacterCount)
		symbols := makeGroupsOfFive(characters)

		// Generate sound samples
		buffer := sample(symbols, config)

		// Play sound to speaker
		trainingStart := time.Now()
		player.Write(buffer) // Blocks until rest of buffer fits into internal buffer (size=0)

		// Print characters to console for comparison
		for i, char := range characters {
			print((strings.ToLower(string(char.character))))
			if i > 0 && (i+1)%5 == 0 {
				print(" ")
			}
		}
		println()

		// Print statistic to console, how often each character appeared
		if config.PrintStatistics {
			charCount := make(map[rune]int)
			for _, char := range characters {
				charCount[char.character]++
			}
			println("Character statistics:")
			for char, count := range charCount {
				println(string(char) + ": " + strconv.Itoa(count))
			}
		}

		// Total training time
		trainingDuration = trainingDuration + time.Since(trainingStart)
		if config.ShowTrainingDuration {
			fmt.Println("Gesamte Übungszeit: " + HumanDuration(trainingDuration))
		}

		// Ask user if he wants another run
		fmt.Print("Nochmal? (j/n): ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		// Read config again for next run - maybe it was adjusted
		if input == "j" {
			config = read_config()
			morse = NewMorseCharacterPool(config)
		}
	}
}
