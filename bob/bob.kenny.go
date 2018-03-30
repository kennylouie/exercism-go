package bob

import (
	"log"
	"regexp"
	"strings"
	"text/scanner"
	"unicode"
)

// classify remarks into different types

// questions end with a question mark after removing white space
func removeWhiteSpace(remark string) string {
	var noWhiteSpace string
	noWhiteSpace = strings.Replace(remark, " ", "", -1)
	return noWhiteSpace
}
func question(remark string) bool {
	var noWhiteSpace string
	noWhiteSpace = removeWhiteSpace(remark)

	var isQuestion bool
	if strings.HasSuffix(noWhiteSpace, "?") {
		isQuestion = true
	} else {
		isQuestion = false
	}
	return isQuestion
}

// yelling remarks have all uppercase alpha characters
func getAlphaChars(remark string) string {
	reg, err := regexp.Compile("[^a-zA-Z]")
	if err != nil {
		log.Fatal(err)
	}
	stringAlpha := reg.ReplaceAllString(remark, "")
	return stringAlpha
}
func getUpperCaseChars(remark string) []string {
	var upperCaseChars []string
	for _, char := range remark {
		if !unicode.IsLower(char) && char != ' ' {
			upperCaseChars = append(upperCaseChars, scanner.TokenString(char))
		}
	}
	return upperCaseChars
}
func yelling(remark string) bool {
	var alphaChars string
	alphaChars = getAlphaChars(remark)

	var lengthAlphaChars int
	lengthAlphaChars = len(alphaChars)

	var upperChars []string
	upperChars = getUpperCaseChars(alphaChars)

	var lengthUpperChars int
	lengthUpperChars = len(upperChars)

	var isYelling bool
	if lengthAlphaChars == lengthUpperChars && lengthAlphaChars != 0 {
		isYelling = true
	} else {
		isYelling = false
	}
	return isYelling
}

// forceful questions have all upper and ends in question mark
func forcefulQuestion(remark string) bool {
	var isForcefulQuestion bool
	if yelling(remark) && question(remark) {
		isForcefulQuestion = true
	} else {
		isForcefulQuestion = false
	}
	return isForcefulQuestion
}

// silence remarks has no alpha nor numeric characters
func getNumChars(remark string) string {
	reg, err := regexp.Compile("[^0-9]")
	if err != nil {
		log.Fatal(err)
	}
	stringNum := reg.ReplaceAllString(remark, "")
	return stringNum
}
func silence(remark string) bool {
	var alphaChars string
	alphaChars = getAlphaChars(remark)

	var lengthAlphaChars int
	lengthAlphaChars = len(alphaChars)

	var numChars string
	numChars = getNumChars(remark)

	var lengthNumChars int
	lengthNumChars = len(numChars)

	var isSilence bool
	if lengthAlphaChars == 0 && lengthNumChars == 0 {
		isSilence = true
	} else {
		isSilence = false
	}
	return isSilence
}

// parsing remark
func Hey(remark string) string {
	var answer string
	if forcefulQuestion(remark) == true {
		answer = "Calm down, I know what I'm doing!"
	} else if question(remark) == true {
		answer = "Sure."
	} else if yelling(remark) == true {
		answer = "Whoa, chill out!"
	} else if silence(remark) == true {
		answer = "Fine. Be that way!"
	} else {
		answer = "Whatever."
	}
	return answer
}
