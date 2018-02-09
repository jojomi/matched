package main

import (
	"fmt"
	"strings"
)

type Match struct {
	Open  string
	Close string
}

func (m *Match) ValidateForString(input string) error {
	// different chars
	if m.Open != m.Close {
		openCount := strings.Count(input, m.Open)
		closeCount := strings.Count(input, m.Close)
		if openCount != closeCount {
			return fmt.Errorf("unbalanced string, checked for %s and %s, opened %d times, closed %d times", m.Open, m.Close, openCount, closeCount)
		}
	}

	// same char (must be even count)
	if m.Open == m.Close {
		count := strings.Count(input, m.Open)
		if count%2 != 0 {
			return fmt.Errorf("unbalanced string, checked for %s, found %d times (odd!)", m.Open, count)
		}
	}

	return nil
}

func splitCharSet(input string) (result []Match) {
	runes := []rune(input)
	var match Match
	for i := 0; i < len(runes)-1; i = i + 2 {
		match = Match{
			Open:  string(runes[i]),
			Close: string(runes[i+1]),
		}
		result = append(result, match)
	}
	return
}
