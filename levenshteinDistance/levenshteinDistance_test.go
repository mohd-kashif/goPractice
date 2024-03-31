package levenshteinDistance

import (
	"testing"
)

func TestGetLavenshteinDistance(t *testing.T) {
	levenshtein := LevenshteinDistance{}

	// Test Case 1: Empty strings
	result := levenshtein.GetLavenshteinDistance("", "")
	if result != 0 {
		t.Errorf("Expected distance for empty strings to be 0, but got %d", result)
	}

	// Test Case 2: One string empty
	result = levenshtein.GetLavenshteinDistance("", "hello")
	if result != 5 {
		t.Errorf("Expected distance for one empty string to be the length of the other string, but got %d", result)
	}

	// Test Case 3: Both strings same
	result = levenshtein.GetLavenshteinDistance("kitten", "kitten")
	if result != 0 {
		t.Errorf("Expected distance for identical strings to be 0, but got %d", result)
	}

	// Test Case 4: Strings of equal length but different characters
	result = levenshtein.GetLavenshteinDistance("kitten", "sitting")
	if result != 3 {
		t.Errorf("Expected distance for strings with different characters to be 3, but got %d", result)
	}

	// Test Case 5: Strings of different lengths
	result = levenshtein.GetLavenshteinDistance("Saturday", "Sunday")
	if result != 3 {
		t.Errorf("Expected distance for strings of different lengths to be the length of the longer string, but got %d", result)
	}

	// Test Case 6: One string is a substring of the other
	result = levenshtein.GetLavenshteinDistance("open", "reopen")
	if result != 2 {
		t.Errorf("Expected distance for one string being a substring of the other to be the difference in lengths, but got %d", result)
	}

	// Test Case 7: Strings with all different characters
	result = levenshtein.GetLavenshteinDistance("abcd", "efgh")
	if result != 4 {
		t.Errorf("Expected distance for strings with all different characters to be the sum of their lengths, but got %d", result)
	}

	// Test Case 8: Strings with all same characters but different lengths
	result = levenshtein.GetLavenshteinDistance("aaaaaa", "aaa")
	if result != 3 {
		t.Errorf("Expected distance for strings with same characters but different lengths to be the difference in lengths, but got %d", result)
	}

	// Test Case 9: Case sensitivity
	result = levenshtein.GetLavenshteinDistance("AbCdEfG", "aBcDeFg")
	if result != 7 {
		t.Errorf("Expected distance for case sensitive strings to be 7, but got %d", result)
	}

	// Test Case 10: Unicode characters
	result = levenshtein.GetLavenshteinDistance("café", "coffee")
	if result != 4 {
		t.Errorf("Expected distance for strings with unicode characters to be 4, but got %d", result)
	}
}
