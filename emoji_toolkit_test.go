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

func TestVariants(t *testing.T) {
	const s = "🌈 The sun ☀️ danced brightly in the sky, illuminating the bustling city 🏙️ filled with laughter 😂 and music 🎶. Children 🎈 played in the park 🌳, while couples ❤️ strolled hand in hand, exchanging sweet nothings 💕. A dog 🐶 chased after a frisbee 🥏, and the smell of delicious food 🍔 wafted from nearby food stalls 🍜. As the afternoon turned to evening 🌅, colorful lights ✨ began to twinkle, setting the stage for a magical night 🌙 filled with dreams 💤 and adventures 🚀!"

	text := ToTextPresentation(s)
	emoji := ToEmojiPresentation(text)

	text2 := ToTextPresentation(emoji)
	emoji2 := ToEmojiPresentation(text2)

	if emoji != emoji2 {
		t.Fatalf("ToEmojiPresentation(%q) = %q; want %s", text2, emoji2, emoji)
	}

	if text != text2 {
		t.Fatalf("ToTextPresentation(%q) = %q; want %s", emoji, text2, text)
	}
}
