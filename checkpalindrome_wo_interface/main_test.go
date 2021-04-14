package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestCheckPalindrome(t *testing.T) {
    assert := assert.New(t)

    var tests = []struct {
        input    string
        expected bool
    }{
        {"a",true},
		{"aaa",true},
		{"aa",true},
		{"abca",false},
		{"AbcBa",true},
		{"@ 2 @",true},
		{"@",true},
		{"@  #",false},
		//{"",false},
		{"Abb a",true},
    }

    for _, test := range tests {
        assert.Equal(CheckPalindrome(test.input), test.expected)
    }
}