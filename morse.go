package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type MorseSymbol int

const (
	Dot MorseSymbol = iota
	Dash
	SymbolSpace
	LetterSpace
	WordSpace
)

type MorseCharacter struct {
	character rune
	symbols   []MorseSymbol
}

type MorseCharacterPool struct {
	symbols []MorseCharacter
}

var morseCharacterTable = []MorseCharacter{
	{'A', []MorseSymbol{Dot, SymbolSpace, Dash}},
	{'B', []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dot}},
	{'C', []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dot}},
	{'D', []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dot}},
	{'E', []MorseSymbol{Dot}},
	{'F', []MorseSymbol{Dot, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dot}},
	{'G', []MorseSymbol{Dash, SymbolSpace, Dash, SymbolSpace, Dot}},
	{'H', []MorseSymbol{Dot, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dot}},
	{'I', []MorseSymbol{Dot, SymbolSpace, Dot}},
	{'J', []MorseSymbol{Dot, SymbolSpace, Dash, SymbolSpace, Dash, SymbolSpace, Dash}},
	{'K', []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dash}},
	{'L', []MorseSymbol{Dot, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dot}},
	{'M', []MorseSymbol{Dash, SymbolSpace, Dash}},
	{'N', []MorseSymbol{Dash, SymbolSpace, Dot}},
	{'O', []MorseSymbol{Dash, SymbolSpace, Dash, SymbolSpace, Dash}},
	{'P', []MorseSymbol{Dot, SymbolSpace, Dash, SymbolSpace, Dash, SymbolSpace, Dot}},
	{'Q', []MorseSymbol{Dash, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dash}},
	{'R', []MorseSymbol{Dot, SymbolSpace, Dash, SymbolSpace, Dot}},
	{'S', []MorseSymbol{Dot, SymbolSpace, Dot, SymbolSpace, Dot}},
	{'T', []MorseSymbol{Dash}},
	{'U', []MorseSymbol{Dot, SymbolSpace, Dot, SymbolSpace, Dash}},
	{'V', []MorseSymbol{Dot, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dash}},
	{'W', []MorseSymbol{Dot, SymbolSpace, Dash, SymbolSpace, Dash}},
	{'X', []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dash}},
	{'Y', []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dash}},
	{'Z', []MorseSymbol{Dash, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dot}},
	{'1', []MorseSymbol{Dot, SymbolSpace, Dash, SymbolSpace, Dash, SymbolSpace, Dash, SymbolSpace, Dash}},
	{'2', []MorseSymbol{Dot, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dash, SymbolSpace, Dash}},
	{'3', []MorseSymbol{Dot, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dash}},
	{'4', []MorseSymbol{Dot, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dash}},
	{'5', []MorseSymbol{Dot, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dot}},
	{'6', []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dot}},
	{'7', []MorseSymbol{Dash, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dot}},
	{'8', []MorseSymbol{Dash, SymbolSpace, Dash, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dot}},
	{'9', []MorseSymbol{Dash, SymbolSpace, Dash, SymbolSpace, Dash, SymbolSpace, Dash, SymbolSpace, Dot}},
	{'0', []MorseSymbol{Dash, SymbolSpace, Dash, SymbolSpace, Dash, SymbolSpace, Dash, SymbolSpace, Dash}},
	{'.', []MorseSymbol{Dot, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dash}},
	{',', []MorseSymbol{Dash, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dash}},
	{';', []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dot}},
	{':', []MorseSymbol{Dash, SymbolSpace, Dash, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dot}},
	{'?', []MorseSymbol{Dot, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dot}},
	{'!', []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dash}},
	{'=', []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dash}},
	{'/', []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dot}},
	{'-', []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dash}},
	{'+', []MorseSymbol{Dot, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dot}},
}

func NewMorseCharacterPool(config Config) *MorseCharacterPool {
	characters := strings.ToUpper(config.Characters)
	pool := &MorseCharacterPool{}
	for _, char := range characters {
		found := false
		for _, morseCharacter := range morseCharacterTable {
			if morseCharacter.character == char {
				pool.symbols = append(pool.symbols, morseCharacter)
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Configured character not found: " + string(char))
		}
	}
	intensiveCharacters := strings.ToUpper(config.IntensiveCharacters)
	for _, char := range intensiveCharacters {
		found := false
		for _, morseCharacter := range morseCharacterTable {
			if morseCharacter.character == char {
				for i := 0; i < config.IntensiveFactor; i++ {
					pool.symbols = append(pool.symbols, morseCharacter)
				}
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Configured intensive character not found: " + string(char))
		}
	}
	return pool
}

func (pool *MorseCharacterPool) GetRandomCharacter() MorseCharacter {
	if len(pool.symbols) == 0 {
		return MorseCharacter{}
	}
	index := rand.Int31n((int32)(len(pool.symbols)))
	return pool.symbols[index]
}

func (pool *MorseCharacterPool) GetRandomCharacters(count int) []MorseCharacter {
	symbols := make([]MorseCharacter, count)
	for i := 0; i < count; i++ {
		symbols[i] = pool.GetRandomCharacter()
	}
	return symbols
}

func makeGroupsOfFive(characters []MorseCharacter) []MorseSymbol {
	var result []MorseSymbol

	// Start sequence: -.-.-
	result = append(result, []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dash, WordSpace}...)

	for i, char := range characters {
		if i > 0 && (i+1)%5 == 0 {
			result = append(result, char.symbols...)
			result = append(result, WordSpace)
		} else {
			result = append(result, char.symbols...)
			result = append(result, LetterSpace)
		}
	}

	// End sequence: .-.-.
	result = append(result, []MorseSymbol{Dot, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace}...)

	return result
}

func getSymbolLengthWithVariant(symbol MorseSymbol, config Config, currentNum int64, maxNum int64, positionStart int64) time.Duration {
	var characterSpacingFactor float64 = 1.0
	var wordSpacingFactor float64 = 1.0

	// Speed curve example: ____..------..___
	lengthTransition := int64(float64(maxNum) * config.VariantTransition / 100.0)
	lengthVariant := int64(float64(maxNum) * config.VariantWidth / 100.0)
	startRising := positionStart
	startVariant := startRising + lengthTransition
	startFalling := startVariant + lengthVariant
	endFalling := startFalling + lengthTransition

	if currentNum > endFalling {
		characterSpacingFactor = config.CharacterSpacingFactor
		wordSpacingFactor = config.WordSpacingFactor
	} else if currentNum >= startFalling {
		characterSpacingFactor = config.CharacterSpacingFactor + (config.VariantCharacterSpacingFactor-config.CharacterSpacingFactor)*((float64(endFalling)-float64(startFalling))-(float64(currentNum)-float64(startFalling)))/(float64(endFalling)-float64(startFalling))
		wordSpacingFactor = config.WordSpacingFactor + (config.VariantWordSpacingFactor-config.WordSpacingFactor)*((float64(endFalling)-float64(startFalling))-(float64(currentNum)-float64(startFalling)))/(float64(endFalling)-float64(startFalling))
	} else if currentNum >= startVariant {
		characterSpacingFactor = config.VariantCharacterSpacingFactor
		wordSpacingFactor = config.VariantWordSpacingFactor
	} else if currentNum >= startRising {
		characterSpacingFactor = config.CharacterSpacingFactor + (config.VariantCharacterSpacingFactor-config.CharacterSpacingFactor)*(float64(currentNum)-float64(startRising))/(float64(startVariant)-float64(startRising))
		wordSpacingFactor = config.WordSpacingFactor + (config.VariantWordSpacingFactor-config.WordSpacingFactor)*(float64(currentNum)-float64(startRising))/(float64(startVariant)-float64(startRising))
	} else {
		characterSpacingFactor = config.CharacterSpacingFactor
		wordSpacingFactor = config.WordSpacingFactor
	}

	return getSymbolLength(symbol, float64(config.WPM), characterSpacingFactor, wordSpacingFactor)
}

func getSymbolLength(symbol MorseSymbol, wpm float64, characterSpacingFactor float64, wordSpacingFactor float64) time.Duration {
	// WPM 60 entspricht 0.1s
	time_dot := time.Duration(float64(time.Second) * 6 / wpm)

	switch symbol {
	case Dot:
		return time_dot
	case Dash:
		return time_dot * 3
	case SymbolSpace:
		return time_dot
	case LetterSpace:
		return time.Duration(float64(time_dot) * 3 * characterSpacingFactor)
	case WordSpace:
		return time.Duration(float64(time_dot) * 7 * wordSpacingFactor)
	default:
		return 0
	}
}
