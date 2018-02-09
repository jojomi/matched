package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateForString(t *testing.T) {
	tests := []struct {
		match Match
		input string
		valid bool
	}{
		{Match{`(`, `)`}, `eval(self)`, true},
		{Match{`(`, `)`}, `(this).eval(((self)))`, true},
		{Match{`(`, `)`}, `eval(self`, false},
		{Match{`(`, `)`}, `evalself)`, false},

		{Match{`"`, `"`}, `"eval"`, true},
		{Match{`"`, `"`}, `"eval" and "go"`, true},
		{Match{`"`, `"`}, `" and "go"`, false},
		{Match{`"`, `"`}, `" and "go`, true},
	}

	for _, tt := range tests {
		err := tt.match.ValidateForString(tt.input)
		assert.Equal(t, tt.valid, err == nil)
	}
}
