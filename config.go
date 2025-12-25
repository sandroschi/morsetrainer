package main

import (
	"io"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	WPM                           int
	Frequency1                    float64
	Frequency2                    float64
	Volume                        float64
	FadeInDuration                float64
	FadeOutDuration               float64
	SampleRate                    float64
	CharacterSpacingFactor        float64
	WordSpacingFactor             float64
	VariantWidth                  float64
	VariantTransition             float64
	VariantCharacterSpacingFactor float64
	VariantWordSpacingFactor      float64
	Characters                    string
	CharacterCount                int
	IntensiveCharacters           string
	IntensiveFactor               int
	PrintStatistics               bool
	ShowTrainingDuration          bool
	DebugWaveForm                 bool
}

func read_config() Config {
	filename := "config.toml"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open file: %s: %v", filename, err)
	}
	defer file.Close()

	var config Config

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	err = toml.Unmarshal(b, &config)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	return config
}
