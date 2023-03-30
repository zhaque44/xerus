package decisionanalysis

import (
	"regexp"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// Customer struct to represent a customer
type Customer struct {
	ID     int
	Name   string
	Email  string
	Churn  bool
	Orders int
}

// Dataset struct to represent a dataset of customers
type Dataset struct {
	Customers []Customer
}

// findPattern searches for a regular expression pattern in a 2D slice of strings
// and returns a 2D slice of matched strings
func findPattern(strings [][]string, pattern string) [][]string {
	matches := [][]string{}
	for _, row := range strings {
		for _, str := range row {
			if ok, _ := regexp.MatchString(pattern, str); ok {
				matches = append(matches, []string{str})
			}
		}
	}
	return matches
}

// classifyChurn uses the findPattern function and a classification algorithm to predict
// whether a customer is likely to churn based on their email address
func ClassifyChurn(customer Customer) bool {
	// Search for a pattern in the email address
	pattern := "gmail"
	matches := findPattern([][]string{{customer.Email}}, pattern)

	// Calculate the Levenshtein distance between the email and the pattern
	var dist int
	if len(matches) > 0 {
		dist = levenshtein.DistanceForStrings([]rune(customer.Email), []rune(matches[0][0]), levenshtein.DefaultOptions)
	} else {
		dist = levenshtein.DistanceForStrings([]rune(customer.Email), []rune(pattern), levenshtein.DefaultOptions)
	}

	// If the distance is less than or equal to half the length of the pattern, predict churn
	if dist <= len(pattern)/2 {
		return true
	} else {
		return false
	}
}
