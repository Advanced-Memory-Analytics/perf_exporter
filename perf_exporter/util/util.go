package util

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func truncateToInt(num int) int {
	// Convert the integer to a string
	numStr := strconv.Itoa(num)

	// Take the first two characters of the string
	truncatedStr := numStr[:2]

	// Convert the truncated string back to an integer
	truncatedInt, err := strconv.Atoi(truncatedStr)
	if err != nil {
		// Handle error if conversion fails
		fmt.Println("Error:", err)
		return 0 // or any default value you prefer
	}

	return truncatedInt
}

func generateRandomPercentages(n int) []float64 {
	// Calculate base percentage and number of values to receive leftover
	basePercentage := 100.0 / float64(n)
	leftoverCount := n / 2

	// Generate base percentages for all values
	percentages := make([]float64, n)
	for i := range percentages {
		percentages[i] = basePercentage
	}

	// Distribute leftover equally among half the values
	leftover := 100.0 - (basePercentage * float64(n))
	for i := 0; i < leftoverCount; i++ {
		percentages[i] += leftover / float64(leftoverCount)
	}

	// Seed the random number generator
	// rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Shuffle the percentages
	// rand.Shuffle(len(percentages), func(i, j int) {
	// 	percentages[i], percentages[j] = percentages[j], percentages[i]
	// })

	// For each percentage, subtract a random float and add it to its symmetrical position on the other half of the slice
	for i := range percentages {
		randomFloat := r.Float64() * percentages[i]
		percentages[i] -= randomFloat
		percentages[(i+n/2)%n] += randomFloat
	}

	sort.SliceStable(percentages, func(i, j int) bool {
		return percentages[i] > percentages[j]
	})

	return percentages
}

func GenerateRandMemLoadString() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	eventSampled := r.Intn(100001-40000) + 40000
	eventCount := int(math.Ceil(float64(eventSampled)/10000) * 10000)

	percentages := generateRandomPercentages(8)

	// Create the template string with placeholders
	memBreakdownHeader := fmt.Sprintf("Samples: %dK of event 'cpu/mem-loads,ldlat=30/P', Event count (approx.): %d\nOverhead\tSamples\tMemory access\n", truncateToInt(eventSampled), eventCount)

	templateMemBreakdown := []string{
		"%.2f%%  \t%d LFB or LFB hit",
		"%.2f%%  \t%d L1 or L1 hit",
		"%.2f%%  \t%d L3 or L3 hit",
		"%.2f%%  \t%d Ldl RAM or RAM hit",
		"%.2f%%  \t%d L2 or L2 hit",
		"%.2f%%  \t%d Uncached or N/A hit",
		"%.2f%%  \t\t%d I/O or N/A hit",
		"%.2f%%  \t\t%d L3 miss",
	}

	// Calculate the total percentage across all categories
	totalPercentage := 0.0

	// Regular expression to match the placeholder indices

	// Process each line
	formattedMemBreakdown := []string{}
	for i, line := range templateMemBreakdown {
		percentage := percentages[i]
		// Replace placeholders with formatted values
		formattedLine := fmt.Sprintf(line, percentage, int(percentage*float64(eventSampled)/100))
		// fmt.Printf("Percentage: %v || eventSampled %v\n", percentage, eventSampled)
		// Update to tal percentage
		totalPercentage += percentage

		formattedMemBreakdown = append(formattedMemBreakdown, formattedLine)
	}

	// Check if total percentage adds up to 100%
	if totalPercentage != 100.00 {
		fmt.Println("Warning: Total percentage does not equal 100% (", totalPercentage, ")")
	}

	// Print the formatted lines as a slice

	return memBreakdownHeader + strings.Join(formattedMemBreakdown, "\n")
}
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
