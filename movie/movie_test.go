package movie

import (
	"testing"
)

func TestMovieMethods(t *testing.T) {
	// Create a new movie
	m := NewMovie("Inception", 2010, "Sci-Fi")

	// Test GetTitle method
	if title := m.GetTitle(); title != "Inception" {
		t.Errorf("Expected title 'Inception', got '%s'", title)
	}

	// Test SetTitle method
	m.SetTitle("Interstellar")
	if title := m.GetTitle(); title != "Interstellar" {
		t.Errorf("Expected title 'Interstellar' after SetTitle, got '%s'", title)
	}

	// Test GetYear method
	if year := m.GetYear(); year != 2010 {
		t.Errorf("Expected year 2010, got %d", year)
	}

	// Test SetYear method
	m.SetYear(2014)
	if year := m.GetYear(); year != 2014 {
		t.Errorf("Expected year 2014 after SetYear, got %d", year)
	}

	// Test GetGenre method
	if genre := m.GetGenre(); genre != "Sci-Fi" {
		t.Errorf("Expected genre 'Sci-Fi', got '%s'", genre)
	}

	// Test SetGenre method
	m.SetGenre("Drama")
	if genre := m.GetGenre(); genre != "Drama" {
		t.Errorf("Expected genre 'Drama' after SetGenre, got '%s'", genre)
	}
}
