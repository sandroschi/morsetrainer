package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

func sample(symbols []MorseSymbol, config Config) []byte {
	// Debug wave form
	var f *os.File
	var w *bufio.Writer
	if config.DebugWaveForm {
		var err error
		f, err = os.Create("/tmp/waveform")
		if err != nil {
			log.Fatalf("Failed to create waveform debug file: %v", err)
		}
		w = bufio.NewWriter(f)
	}
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	dotWritten := false

	// Get total length in microseconds
	var totalLength int64 = 0
	for _, symbol := range symbols {
		symbolLength := getSymbolLength(symbol, config)
		totalLength += symbolLength.Microseconds()
	}
	// Total length in samples
	totalLength = (totalLength * (int64(config.SampleRate))) / 1000000
	buffer := make([]byte, 2*totalLength) // 2 Bytes per sample
	var bufferIndex int64 = 0

	// Write all samples to buffer
	for _, symbol := range symbols {
		symbolLength := getSymbolLength(symbol, config)
		samplesToWrite := int(config.SampleRate * float64(symbolLength) / float64(time.Second))
		t := 0.0
		sample := 0.0

		if symbol == Dot || symbol == Dash {
			// Fading sample count
			fadeInSamples := int(float64(config.SampleRate) / float64(time.Second) * float64(config.FadeInDuration) * float64(time.Millisecond))
			fadeOutSamples := int(float64(config.SampleRate) / float64(time.Second) * float64(config.FadeOutDuration) * float64(time.Millisecond))

			// Frequency
			var frequency float64
			if symbol == Dot {
				frequency = config.Frequency1
			} else {
				frequency = config.Frequency2
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

				// Store in buffer
				v := int16(sample * 32767)
				buffer[bufferIndex] = byte(v)
				buffer[bufferIndex+1] = byte(v >> 8)
				bufferIndex += 2
				t += 1.0 / config.SampleRate

				// Debug wave form
				if symbol == Dot && !dotWritten && config.DebugWaveForm {
					_, _ = fmt.Fprintf(w, "%d\n", v)
				}
			}
			if symbol == Dot {
				dotWritten = true
			}
		} else { // Spaces are zeros in the buffer. If you stop playing and sleep instead you have to deal with clipping.
			// Generate silence samples
			for i := 0; i < samplesToWrite; i++ {
				buffer[bufferIndex] = 0
				buffer[bufferIndex+1] = 0
				bufferIndex += 2
			}
		}
	}

	if config.DebugWaveForm {
		w.Flush()
	}

	return buffer
}
