package main

import (
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
	{'?', []MorseSymbol{Dot, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dot}},
	{'=', []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dot, SymbolSpace, Dash}},
}

func NewMorseCharacterPool(config Config) *MorseCharacterPool {
	characters := strings.ToUpper(config.Characters)
	pool := &MorseCharacterPool{}
	for _, char := range characters {
		for _, morseCharacter := range morseCharacterTable {
			if morseCharacter.character == char {
				pool.symbols = append(pool.symbols, morseCharacter)
				break
			}
		}
	}
	for _, char := range config.IntensiveCharacters {
		for _, morseCharacter := range morseCharacterTable {
			if morseCharacter.character == char {
				for i := 0; i < config.IntensiveFactor; i++ {
					pool.symbols = append(pool.symbols, morseCharacter)
				}
				break
			}
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
	result = append(result, []MorseSymbol{Dash, SymbolSpace, Dot, SymbolSpace, Dash, SymbolSpace, Dot, SymbolSpace, Dash, WordSpace, WordSpace}...)

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

func getSymbolLength(symbol MorseSymbol, config Config) time.Duration {
	// WPM 60 entspricht 0.1s
	time_dot := time.Duration(float64(time.Second) * 6 / float64(config.WPM))

	switch symbol {
	case Dot:
		return time_dot
	case Dash:
		return time_dot * 3
	case SymbolSpace:
		return time_dot
	case LetterSpace:
		return time.Duration(float64(time_dot) * 3 * config.CharacterSpacingFactor)
	case WordSpace:
		return time.Duration(float64(time_dot) * 7 * config.WordSpacingFactor)
	default:
		return 0
	}
}
