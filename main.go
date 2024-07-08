package main

import (
	
	"os"
)

// split function takes a string and returns a slice of strings
// It splits the input string using the provided separator (default is " ")
// It removes any empty strings from the result
func split(str string) []string {
	var result []string
	var finalResult []string
	start := 0
	sep := " "

	for i := 0; i < len(str); i++ {
		if i+len(sep) <= len(str) && str[i:i+len(sep)] == sep {
			result = append(result, str[start:i])
			start = i + len(sep)
			i += len(sep) - 1
		}
	}
	result = append(result, str[start:])
	for _, v := range result {
		if v != "" {
			finalResult = append(finalResult, v)
		}
	}
	return finalResult
}

// rotateByOne function takes a slice of strings and rotates the elements by one position
// The first element becomes the last element, and all other elements are shifted to the left
func rotateByOne(words []string) []string {
	rotated := make([]string, len(words))
	for i := 1; i < len(words); i++ {
		rotated[i-1] = words[i]
	}
	rotated[(len(words) - 1)] = words[0]
	return rotated
}

// joinWords function takes a slice of strings and joins them into a single string
// It adds a space between each word
func joinWords(words []string) string {
	result := ""
	for i, word := range words {
		if i > 0 {
			result += " "
		}
		result += word
	}
	return result
}

func main() {
	// Check if the correct number of command-line arguments are provided
	if len(os.Args) != 2 {
		os.Stdout.WriteString("\n")
		return
	}
	args := os.Args[1]

	// Split the input string into words
	splat := split(args)
	// Rotate the words by one position
	rotated := rotateByOne(splat)
	// Join the rotated words back into a string
	joined := joinWords(rotated)

	// Write the result to stdout
	os.Stdout.WriteString(joined + "\n")
}
