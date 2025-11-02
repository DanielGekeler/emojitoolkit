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

func FuzzContainsEmoji(f *testing.F) {
	f.Add("A")
	f.Add("1")
	f.Add("⏳")
	f.Add("🌍")
	f.Add("☀")
	f.Add("♻")
	f.Add("☀️")
	f.Add("♻️")
	f.Add("1️⃣")

	f.Fuzz(func(t *testing.T, s string) {
		ContainsEmoji(s)
	})
}

func TestData(t *testing.T) {
	ranges := [][]int32{
		emoji_ranges1,
		emoji_ranges2,
		emoji_ranges3,
		variant_ranges,
	}

	for _, rs := range ranges {
		if len(rs)%2 != 0 {
			t.Fail()
		}
	}
}

func TestToTextPresentation(t *testing.T) {
	testCases := map[string]string{
		"":     "",
		"A":    "A",
		"1":    "1",
		"1️⃣":  "1",
		"1️⃣.": "1.",
		"⏳":    "⏳\uFE0E",
		"🌍":    "🌍\uFE0E",
		"☀":    "☀\uFE0E",
		"♻":    "♻\uFE0E",
		"☀️":   "☀\uFE0E",
		"♻️":   "♻\uFE0E",

		"⏳.":  "⏳\uFE0E.",
		"🌍.":  "🌍\uFE0E.",
		".⏳.": ".⏳\uFE0E.",
		".🌍.": ".🌍\uFE0E.",

		".🌍.🌍..🌍.": ".🌍\uFE0E.🌍\uFE0E..🌍\uFE0E.",
	}

	for input, expected := range testCases {
		result := ToTextPresentation(input)
		if result != expected {
			t.Fatalf("ToTextPresentation(\"%s\") = %s; want %s", input, result, expected)
		}
	}
}

func TestToEmojiPresentation(t *testing.T) {
	testCases := map[string]string{
		"":     "",
		"A":    "A",
		"1":    "1",
		"1️⃣":  "1️⃣",
		"1️⃣.": "1️⃣.",
		"⏳":    "⏳\uFE0F",
		"🌍":    "🌍\uFE0F",

		"⏳\uFE0E": "⏳\uFE0F",
		"🌍\uFE0E": "🌍\uFE0F",
		"☀\uFE0E": "☀️",
		"♻\uFE0E": "♻️",
		"☀":       "☀️",
		"♻":       "♻️",

		"⏳\uFE0E.":  "⏳\uFE0F.",
		"🌍\uFE0E.":  "🌍\uFE0F.",
		".⏳\uFE0E.": ".⏳\uFE0F.",
		".🌍\uFE0E.": ".🌍\uFE0F.",

		".🌍\uFE0E.🌍..🌍\uFE0E.": ".🌍\uFE0F.🌍\uFE0F..🌍\uFE0F.",
	}

	for input, expected := range testCases {
		result := ToEmojiPresentation(input)
		t.Logf("%q -> %q", input, result)
		if result != expected {
			t.Fatalf("ToEmojiPresentation(\"%s\") = %s; want %s", input, result, expected)
		}
	}
}
