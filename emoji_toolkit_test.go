package emojitoolkit

import (
	"testing"
)

func TestIsSingleCharacterEmoji(t *testing.T) {
	testCases := map[rune]bool{
		'A': false,
		'1': false,
		'⏳': true,
		'🌍': true,
		'☀': false,
		'♻': false,
	}

	for input, expected := range testCases {
		result := IsSingleCharacterEmoji(input)
		if result != expected {
			t.Fatalf("IsSingleCharacterEmoji('%c') = %v; want %v", input, result, expected)
		}
	}
}

func TestContainsEmoji(t *testing.T) {
	testCases := map[string]bool{
		"A":   false,
		"1":   false,
		"⏳":   true,
		"🌍":   true,
		"☀":   false,
		"♻":   false,
		"☀️":  true,
		"♻️":  true,
		"1️⃣": true,

		"⏳.": true,
		"🌍.": true,
		"☀.": false,
		"♻.": false,
	}

	for input, expected := range testCases {
		result := ContainsEmoji(input)
		if result != expected {
			t.Fatalf("ContainsEmoji(\"%s\") = %v; want %v", input, result, expected)
		}
	}
}
