package test

import (
	"fmt"
	"testing"

	"github.com/vaibhavp1964/go-redis/parser"
)

const (
	testString_correct_echo_foo = "*2\r\n$4\r\nECHO\r\n$3\r\nFOO\r\n"
)

func TestParse(t *testing.T) {
	expectedTokens := []string{"ECHO", "FOO"}
	receivedTokens, _, _ := parser.ParseInput(testString_correct_echo_foo, 0)

	fmt.Println("Received tokens: ", receivedTokens)

	if len(expectedTokens) != len(receivedTokens) {
		t.Errorf("error while parsing test string: %s", "incorrect tokenisation")
	}

	for index := range expectedTokens {
		if expectedTokens[index] != receivedTokens[index] {
			t.Errorf("error while parsing test string: %s", "incorrect tokenisation")
		}
	}
}
