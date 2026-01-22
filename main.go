package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Print("Enter a cooking time: \n")

	quantity, captureErr := captureInput()
	if captureErr != nil {
		fmt.Printf("problem trying to read input: %s\n", captureErr.Error())
		os.Exit(1)
	}

	// check whether the input makes sense as time
	// and if not treat is a raw quantity of seconds
	var treatAsTotalSeconds bool
	for idx, char := range quantity {
		if char >= six && idx != len(quantity)-1 {
			treatAsTotalSeconds = true
		}
	}

	var converted int
	var err error
	if treatAsTotalSeconds {
		converted, err = strconv.Atoi(quantity)
		if err != nil {
			fmt.Printf("problem converting sepcified quantity: %s", err.Error())
			os.Exit(1)
		}
	} else {
		minsAndSecs, err := deriveMinutesAndSeconds(quantity)
		if err != nil {
			fmt.Printf("problem trying to normalize entered amount: %s\n", err.Error())
			os.Exit(1)
		}
		converted = (minsAndSecs.minutes * 60) + minsAndSecs.seconds
	}

	if err := cook(converted); err != nil {
		fmt.Printf("problem cooking: %s\n", err.Error())
		os.Exit(1)
	}
}

const colon = rune(32)
const zero = rune(48)
const six = rune(54)
const nine = rune(57)

// Read input from STDIN and validate the string once captured.
// Trying to validate a single character at time means relying on
// external libraries which I wanted to avoid
func captureInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	var validated string
	var nonZeroSeen bool

	// filter out invalid characters
	for idx, char := range input {
		// skip colon characters
		if char == colon {
			continue
		}

		// ignore all but numeric characters
		if char < zero || char > nine {
			continue
		}

		// ignore leading zeros but make sure to keep the rest
		if char == zero {
			if idx == 0 || !nonZeroSeen {
				// fmt.Printf("pos is zero: %d\n", idx)
				continue
			}
		} else {
			nonZeroSeen = true
		}

		if nonZeroSeen {
			validated += string(char)
		}
	}
	// fmt.Printf("validated: [%v]\n", validated)

	// cap length at 4 numeric characters to allow for 99:99
	if len(validated) > 4 {
		return "", fmt.Errorf("max input is 99:99, got length of %d", len(validated))
	}

	if validated == "" {
		return "", fmt.Errorf("0 is not a valid length of time")
	}

	return validated, nil
}

type MinsAndSecs struct {
	minutes int
	seconds int
}

func deriveMinutesAndSeconds(quantity string) (MinsAndSecs, error) {
	if len(quantity) > 2 {
		seconds := quantity[1:]
		secsAsInt, err := strconv.Atoi(seconds)
		if err != nil {
			return MinsAndSecs{}, fmt.Errorf("error converting seconds: %w", err)
		}
		// fmt.Printf("secsAsInt?: %d\n", secsAsInt)

		minutes := quantity[:1]
		minsAsInt, err := strconv.Atoi(minutes)
		if err != nil {
			return MinsAndSecs{}, fmt.Errorf("error converting minutes: %w", err)
		}

		minsAndSecs := MinsAndSecs{
			minutes: minsAsInt,
			seconds: secsAsInt,
		}
		return minsAndSecs, nil
	}

	secsAsInt, err := strconv.Atoi(quantity)
	if err != nil {
		return MinsAndSecs{}, fmt.Errorf("error converting quantity: %w", err)
	}

	return MinsAndSecs{seconds: secsAsInt}, nil
}

// perform the cooking countdown
func cook(cookTime int) error {
	fmt.Println("cooking...")
	timeLeft := cookTime
	for {
		minutes := (timeLeft / 60)
		// fmt.Printf("minutes: %d\n", minutes)

		seconds := timeLeft - (minutes * 60)
		// fmt.Printf("seconds: %d\n", seconds)

		fmt.Printf("%02d:%02d\n", minutes, seconds)

		time.Sleep(time.Second * 1)
		if timeLeft == 0 {
			break
		}
		timeLeft--
	}
	return nil
}
