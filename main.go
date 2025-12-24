package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/oto"
)

func main() {
	var config = read_config()
	var sampleRate = config.SampleRate
	var morse = NewMorseCharacterPool(config)

	ctx, _ := oto.NewContext((int)(sampleRate), 1, 2, 0)
	player := ctx.NewPlayer()
	defer player.Close()

	reader := bufio.NewReader(os.Stdin)
	var input string = "j"
	var trainingDuration time.Duration = 0

	var t float64
	fadeOutDuration := time.Duration(float64(time.Millisecond) * config.FadeOutDuration)
	fadeInDuration := time.Duration(float64(time.Millisecond) * config.FadeInDuration)
	fadeOutDeadSamples := config.FadeOutDeadSamples
	fadeInDeadSamples := config.FadeInDeadSamples

	for input == "j" {
		characters := morse.GetRandomCharacters(config.CharacterCount)
		symbols := makeGroupsOfFive(characters)
		trainingStart := time.Now()

		// Debug clipping
		// f, _ := os.Create("/tmp/yourfile")
		// defer f.Close()
		// w := bufio.NewWriter(f)

		for _, symbol := range symbols {
			symbolLength := getSymbolLength(symbol, config)

			t = 0.0
			sample := 0.0

			for start := time.Now(); time.Since(start) < symbolLength; {
				if symbol == Dot || symbol == Dash {
					samplesToWrite := int(sampleRate * float64(symbolLength) / float64(time.Second))
					buf := make([]byte, samplesToWrite*2+fadeInDeadSamples*2+fadeOutDeadSamples*2) // A buffer to hold all samples for this symbol - 2 bytes per sample
					fadeInSamples := int(float64(sampleRate) * float64(fadeInDuration) / float64(time.Second))
					fadeOutSamples := int(float64(sampleRate) * float64(fadeOutDuration) / float64(time.Second))
					var frequency float64
					if symbol == Dot {
						frequency = config.Frequency1
					} else {
						frequency = config.Frequency2
					}

					// Dead samples help with clipping
					for i := 0; i < fadeInDeadSamples; i++ {
						buf[2*i] = 0
						buf[2*i+1] = 0
					}

					// Generate samples
					for i := 0; i < samplesToWrite; i++ {
						// Fading effect
						fading := 1.0
						if i < fadeInSamples {
							fading = float64(i) / float64(fadeInSamples)
						} else if i >= samplesToWrite-fadeOutSamples {
							fading = float64(samplesToWrite-i) / float64(fadeOutSamples)
						}

						// Generate sample
						sample = math.Sin(2*math.Pi*frequency*t) * config.Volume * fading

						v := int16(sample * 32767)
						//_, _ = fmt.Fprintf(w, "%d\n", v)
						buf[fadeInDeadSamples+2*i] = byte(v)
						buf[fadeInDeadSamples+2*i+1] = byte(v >> 8)
						t += 1.0 / sampleRate
					}

					// Dead samples help with clipping
					for i := 0; i < fadeOutDeadSamples; i++ {
						buf[fadeInDeadSamples*2+samplesToWrite*2+i*2] = 0
						buf[fadeInDeadSamples*2+samplesToWrite*2+i*2+1] = 0
					}
					player.Write(buf) // Blocks until rest of buffer fits into internal buffer
					break
				} else {
					time.Sleep(symbolLength)
					break
				}
			}

			// //println(symbolLength) // Dash = 225ms, Dot = 75ms
			// fmt.Fprintf(w, "Ende")
			// w.Flush()
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

		// Total training time
		trainingDuration = trainingDuration + time.Since(trainingStart)
		if config.ShowTrainingDuration {
			fmt.Println("Gesamte Übungszeit: " + HumanDuration(trainingDuration))
		}

		// Ask user if he wants another run
		fmt.Print("Nochmal? (j/n): ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
	}
}
