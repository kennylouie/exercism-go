package bob

import (
	"strings"
	"unicode"
)

func Hey(input string) string {

	/*
	   Bob answers 'Sure.' if you ask him a question.
	   He answers 'Whoa, chill out!' if you yell at him.
	   He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
	   He says 'Fine. Be that way!' if you address him without actually saying anything.
	*/

	// check if value is a string
	if input == "" {
		return "Fine. Be that way!"
	}

	FirstCharByte := input[0]
	LastCharByte := input[len(input)-1]
	AllCapsInput := strings.ToUpper(input)

	if strings.Contains(input, "\t") == true || (string(FirstCharByte) == " " && string(LastCharByte) == " ") {
		return "Fine. Be that way!"
	}

	// check for ? ending + all caps + first letter is a letter
	if input == AllCapsInput && strings.HasSuffix(input, "?") == true && unicode.IsLetter(rune(FirstCharByte)) == true {
		return "Calm down, I know what I'm doing!"
	}

	// check for ? ending + no caps + last letter is space
	if strings.HasSuffix(input, "?") == true || (strings.Contains(input, "?") && string(LastCharByte) == " ") {
		return "Sure."
	}

	// check for all caps + ! ending
	if input == AllCapsInput && (strings.HasSuffix(input, "!") || unicode.IsLetter(rune(FirstCharByte)) == true) {
		return "Whoa, chill out!"
	}

	// default
	return "Whatever."
}
