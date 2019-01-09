package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/sethvargo/go-password/password"
)

var logger *log.Logger

var verboseLogging = false

func main() {
	logger = log.New(os.Stderr, "", 0)

	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	setVerboseLogging(args)

	length := getLength(args)
	logf("Password will be %d characters", length)

	split, allowRepeatCharacters := getSplit(length, args)
	logf("Password will have %d digits and %d symbols", split[0], split[1])

	pw, err := password.Generate(length, split[0], split[1], false, allowRepeatCharacters)
	if err != nil {
		logf("Password failed to generate: %s", err.Error())
		os.Exit(1)
	}

	logf("Password generated")

	fmt.Print(pw)

	os.Exit(0)
}

func logf(msg string, data ...interface{}) {
	if verboseLogging {
		logger.Printf(msg, data...)
	}
}

func setVerboseLogging(args []string) {
	verboseLogging = false

	for _, arg := range args {
		if arg == "-v" {
			verboseLogging = true
			break
		}
	}
}

func getLength(args []string) int {
	minLength := 8
	maxLength := 4096

	length := 32

	for _, arg := range args {
		if len(arg) > 9 && arg[0:9] == "--length=" {
			parts := strings.Split(arg, "=")
			if len(parts) == 2 {
				requestedLength, err := strconv.Atoi(parts[1])
				if err == nil {
					if requestedLength >= minLength && requestedLength <= maxLength {
						length = requestedLength
					} else if requestedLength > maxLength {
						length = maxLength
						logf("Specified password length too long; overriding to %d characters", maxLength)
					} else if requestedLength < minLength {
						length = minLength
						logf("Specified password length too short; overriding to %d characters", minLength)
					}
				}
			}

			break
		}
	}

	return length
}

func getSplit(length int, args []string) ([2]int, bool) {
	allowRepeatCharacters := false
	if length > 26 {
		allowRepeatCharacters = true
		logf("Password will include repeated characters")
	} else {
		logf("Password will not include repeated characters")
	}

	includeSymbols := true
	for _, arg := range args {
		if arg == "--no-symbols" {
			includeSymbols = false
			break
		}
	}

	modifier := float64(2)
	if includeSymbols {
		modifier = 3
	}

	split := math.Floor(float64(length) / modifier)

	numberofDigits := int(split)
	if numberofDigits > 10 && !allowRepeatCharacters {
		numberofDigits = 10
	}

	numberOfSymbols := int(split)
	if !includeSymbols {
		numberOfSymbols = 0
	}

	return [2]int{numberofDigits, numberOfSymbols}, allowRepeatCharacters
}
