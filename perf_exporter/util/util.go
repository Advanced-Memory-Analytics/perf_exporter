package util

import (
	"fmt"
	"regexp"
	"strconv"
)

func ConvertNumberFormat(value string) (int, error) {
	// Regular expression to match digits and optional "k" suffix
	re := regexp.MustCompile(`^(\d+)([kK]?)$`)

	// Capture matched groups
	matches := re.FindStringSubmatch(value)
	if matches == nil {
		return -1, fmt.Errorf("unable to parse")
	}

	numberStr, suffix := matches[1], matches[2]

	// Convert string to int
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return -1, fmt.Errorf("unable to convert string to int: %v", err)
	}

	// Convert based on suffix
	switch suffix {
	case "":
		return (number), nil
	case "k":
		return number * 1000, nil
	case "K":
		return number * 1000, nil
	case "m":
		return number * 1000000, nil
	default:
		return -1, (fmt.Errorf("unexpected integer suffix: %s", suffix)) // Handle unexpected cases
	}
}
