package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vaibhavp1964/go-redis/command"
	"github.com/vaibhavp1964/go-redis/constants"
)

func Parse(input string) command.Command {
	return command.Command{}
}

func ParseInput(input string, pos int) ([]string, int, bool) {
	inputLength := len(input[pos:])
	index := pos
	tokens := make([]string, 0)

	for ; index < inputLength; index++ {
		newTokens := make([]string, 0)
		valid := false
		newPos := 0

		leadingChar := string(input[index])
		EOL := false

		switch leadingChar {
		case string(constants.CRLF_LEADING):
			// case for CRLF
			if encounteredCRLF(input, index) {
				newPos, valid = jumpOverCRLF(input, index)
				if newPos == -2 {
					EOL = true
				}
			}
			break
		case string(constants.ARRAY):
			// case for array
			newTokens, newPos, valid = getArray(input, index)
			break
		case string(constants.BULK_STRING):
			// case for strings
			newTokens, newPos, valid = getString(input, index)
			break
		default:
			fmt.Println(leadingChar)
		}

		if EOL {
			break
		}
		if !valid {
			return []string{}, -1, false
		}
		index = newPos - 1
		tokens = append(tokens, newTokens...)
	}

	return tokens, index, true
}

func getLength(input string, pos int) (int, int, bool) {
	length := 0
	index := pos + 1

	for ; index < len(input); index++ {
		digit, err := strconv.Atoi(string(input[index]))
		if encounteredCRLF(input, index) {
			break
		}
		if err != nil {
			return -1, -1, false
		}

		length *= 10
		length += digit
	}

	jump, valid := jumpOverCRLF(input, index)
	if !valid {
		return -1, -1, false
	}

	return length, jump, true
}

func getString(input string, pos int) ([]string, int, bool) {
	// get length of strings
	strLength, index, valid := getLength(input, pos)
	if !valid {
		return []string{}, -1, false
	}

	inputString := make([]string, 0)

	for ; index < len(input); index++ {
		if encounteredCRLF(input, index) {
			break
		}
		inputString = append(inputString, string(input[index]))
	}

	newString := strings.Join(inputString, "")

	// validate length of string
	if len(newString) != strLength {
		return []string{}, -1, false
	}

	return []string{strings.Join(inputString, "")}, index, true
}

func getArray(input string, pos int) ([]string, int, bool) {
	// get length of array
	arrayLength, index, valid := getLength(input, pos)
	if !valid {
		return []string{}, -1, false
	}

	// recursively parse for array elements
	array, newPos, valid := ParseInput(input, index)

	// validate correctness
	if !valid || arrayLength != len(array) {
		return []string{}, -1, false
	}

	return array, newPos, true
}

func encounteredCRLF(input string, pos int) bool {
	chars := make([]string, 0)
	index := pos

	for ; index < pos+len(constants.CRLF); index++ {
		chars = append(chars, string(input[index]))
	}

	return strings.Compare(strings.Join(chars, ""), string(constants.CRLF)) == 0
}

func jumpOverCRLF(input string, pos int) (int, bool) {
	isCRLF := encounteredCRLF(input, pos)
	if !isCRLF {
		return -1, false
	}
	newPos := pos + len(string(constants.CRLF))
	if newPos >= len(input) {
		return -2, false
	}

	return newPos, true
}
