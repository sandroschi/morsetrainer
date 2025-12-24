package main

import (
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	WPM                    int
	Frequency1             float64
	Frequency2             float64
	Volume                 float64
	FadeInDuration         float64
	FadeOutDuration        float64
	FadeInDeadSamples      int
	FadeOutDeadSamples     int
	SampleRate             float64
	CharacterSpacingFactor float64
	WordSpacingFactor      float64
	Characters             string
	CharacterCount         int
	IntensiveCharacters    string
	IntensiveFactor        int
	PrintStatistics        bool
	ShowTrainingDuration   bool
}

func read_config() Config {
	file, err := os.Open("config.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var config Config

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}

	return config
}
