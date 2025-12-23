package main

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/oto"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	const sampleRate = 44100
	var config = read_config()
	var morse = NewMorseCharacterPool(config)

	ctx, _ := oto.NewContext(sampleRate, 1, 2, 0)
	player := ctx.NewPlayer()
	defer player.Close()
	var t float64
	buf := make([]byte, 1024)
	fadeOutDuration := time.Duration(float64(time.Millisecond) * 20)

	characters := morse.GetRandomCharacters(config.CharacterCount)
	symbols := makeGroupsOfFive(characters)

	for _, symbol := range symbols {
		symbolLength := getSymbolLength(symbol, config)
		//symbolEnd := time.Now().Add(symbolLength)

		// for time.Now().Before(symbolEnd) {
		// 	for i := 0; i < len(buf)/2; i++ {
		// 		var sample float64
		// 		if symbol == Dot || symbol == Dash {
		// 			sample = (math.Sin(2*math.Pi*config.Frequency1*t) +
		// 				math.Sin(2*math.Pi*config.Frequency2*t)) * 0.5
		// 		} else {
		// 			sample = 0
		// 		}

		// 		v := int16(sample * 32767)
		// 		buf[2*i] = byte(v)
		// 		buf[2*i+1] = byte(v >> 8)
		// 		t += 1.0 / sampleRate
		// 	}
		// 	player.Write(buf)
		// }
		t = 0
		for start := time.Now(); time.Since(start) < symbolLength; {
			if symbol == Dot || symbol == Dash {
				timeRemaining := symbolLength - time.Since(start)

				if timeRemaining < fadeOutDuration {
					// Apply fade out
					fadeOutFactor := float64(timeRemaining) / float64(fadeOutDuration)
					for i := 0; i < len(buf)/2; i++ {
						sample := (math.Sin(2*math.Pi*config.Frequency1*t) +
							math.Sin(2*math.Pi*config.Frequency2*t)) * 0.5 * fadeOutFactor

						v := int16(sample * 32767)
						buf[2*i] = byte(v)
						buf[2*i+1] = byte(v >> 8)
						t += 1.0 / sampleRate
					}
					player.Write(buf)
				} else {
					// Normal sound
					for i := 0; i < len(buf)/2; i++ {
						sample := (math.Sin(2*math.Pi*config.Frequency1*t) +
							math.Sin(2*math.Pi*config.Frequency2*t)) * 0.5

						v := int16(sample * 32767)
						buf[2*i] = byte(v)
						buf[2*i+1] = byte(v >> 8)
						t += 1.0 / sampleRate
					}
					player.Write(buf)
				}

				// for i := 0; i < len(buf)/2; i++ {
				// 	sample := (math.Sin(2*math.Pi*config.Frequency1*t) +
				// 		math.Sin(2*math.Pi*config.Frequency2*t)) * 0.5

				// 	v := int16(sample * 32767)
				// 	buf[2*i] = byte(v)
				// 	buf[2*i+1] = byte(v >> 8)
				// 	t += 1.0 / sampleRate
				// }
				// player.Write(buf)
			}
		}
	}

	// Print characters to console
	for i, char := range characters {
		// Print character as lowercase
		print((strings.ToLower(string(char.character))))
		if i > 0 && (i+1)%5 == 0 {
			print(" ")
		}
	}
	println()

	// Print statistic to console how often each character appeared
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
}
