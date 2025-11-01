package emojitoolkit

import (
	"slices"
)

// Supported Unicode version
const Version = "17.0.0"

//go:generate go run generator/main.go

// Matches codepoints that are default emoji presentation character ([ED-6])
// that can be repressented as a single rune / codepoint.
//
// Codepoints that appear as text by default ([ED-7]) but can appear as emoji
// using emoji presentation sequence ([ED-9a]) are not matched because
// the pressence of U+FE0F VARIATION SELECTOR-16 (VS16) can not be verified from a single rune.
//
// See section Basic_Emoji of [emoji-sequences.txt] for a list of matching emojis.
//
// # Examples:
//
//	'A' -> false // U+0041
//	'⏳' -> true // U+231B
//	'🌍' -> true // U+1F30D
//	'☀' -> false // U+2600
//	'♻' -> false // U+267B
//
// This function can also be defined as Emoji_Presentation=Yes and Emoji_Component=No
//
// [ED-6]: https://www.unicode.org/reports/tr51/#def_emoji_presentation
// [ED-7]: https://www.unicode.org/reports/tr51/#def_text_presentation
// [ED-9a]: https://www.unicode.org/reports/tr51/#def_emoji_presentation_sequence
// [emoji-sequences.txt]: https://www.unicode.org/Public/emoji/latest/emoji-sequences.txt
func IsSingleCharacterEmoji(r rune) bool {
	return isInRange(r, emoji_ranges1)
}

// Matches all* emojis
//   - default emoji presentation character ([ED-6])
//   - emoji presentation sequence ([ED-9a])
//   - emoji keycap sequence ([ED-14c])
//   - emoji flag sequence ([ED-14])
//
// This should include all basic emoji defined by [ED-20].
//
// See [emoji-sequences.txt] for a list of matching emojis.
//
// # Examples:
//
//	"A" -> false // U+0041
//	"⏳" -> true // U+231B
//	"🌍" -> true // U+1F30D
//	"☀" -> false // U+2600
//	"♻" -> false // U+267B
//	"☀️":  true // U+2600 U+FE0F
//	"♻️":  true // U+267B U+FE0F
//
// [ED-6]: https://www.unicode.org/reports/tr51/#def_emoji_presentation
// [ED-9a]: https://www.unicode.org/reports/tr51/#def_emoji_presentation_sequence
// [ED-14]: https://www.unicode.org/reports/tr51/#def_emoji_flag_sequence
// [ED-14c]: https://www.unicode.org/reports/tr51/#def_emoji_keycap_sequence
// [ED-20]: https://www.unicode.org/reports/tr51/#def_basic_emoji_set
// [emoji-sequences.txt]: https://www.unicode.org/Public/emoji/latest/emoji-sequences.txt
func ContainsEmoji(s string) bool {
	runes := []rune(s)
	l := len(runes)
	for i, r := range runes {
		if IsSingleCharacterEmoji(r) {
			return true
		}

		if i+1 >= l {
			break // skip last rune because there is no next rune
		}

		next := runes[i+1]
		if isInRange(r, emoji_ranges2) && next == vs16 {
			// ED-7 default text presentation character
			return true
		}

		if isInRange(r, emoji_ranges3) && (next == vs16 || (next >= light_skin && next <= dark_skin)) {
			// ED-22 RGI emoji modifier sequence set
			return true
		}

		if IsFlagSequence(runes[i:]) {
			return true
		}
	}
	return false
}

// Check if a rune is within any range of a set of ranges.
// These sets are an array of int32 with each range being defined
// as two consecutive int32 defining the first and last element
// of a range respectively.
func isInRange(r rune, ranges []int32) bool {
	for i := 0; i < len(ranges); i += 2 {
		if r >= rune(ranges[i]) && r <= rune(ranges[i+1]) {
			return true
		}
	}
	return false
}

func isInRangeF(ranges []int32) func(rune) bool {
	return func(r rune) bool { return isInRange(r, ranges) }
}

// Matches flag emojis officially known as emoji flag sequence ([ED-14]).
// Does not check if the flag is valid.
//
// [ED-14]: https://www.unicode.org/reports/tr51/#def_emoji_flag_sequence
func IsFlagSequence(runes []rune) bool {
	if len(runes) < 2 {
		return false
	}

	a := runes[0]
	b := runes[1]
	return a >= flagA && a <= flagB && b >= flagA && b <= flagB
}

// Matches a string that contains atleast one flag emoji officially known as emoji flag sequence ([ED-14]).
// Does not check if the flag is valid.
//
// [ED-14]: https://www.unicode.org/reports/tr51/#def_emoji_flag_sequence
func ContainsFlag(s string) bool {
	runes := []rune(s)

	for i := 0; i < len(runes)-1; i++ {
		if IsFlagSequence(runes[i:]) {
			return true
		}
	}
	return false
}

func ToTextPresentation(s string) string {
	if s == "" {
		return s
	}

	// returns the adjusted string and how many were taken
	f := func(rs []rune) ([]rune, int) {
		i := slices.IndexFunc(rs, isInRangeF(variant_ranges))
		if i < 0 {
			return rs, len(rs)
		}
		x := i + 1

		if i+1 < len(rs) && rs[i+1] == vs16 {
			x++
		}

		// Special treatment for 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, #, *
		if r := rs[i]; (r >= '0' && r <= '9') || r == '#' || r == '*' {
			// ED-14c emoji keycap sequence
			if i+2 < len(rs) && rs[i+2] == keycap {
				x++
			}
			return rs[:i+1], x
		}

		ret := make([]rune, i+1)
		copy(ret, rs[:i+1])
		return append(ret, vs15), x
	}

	runes := []rune(s)
	ret := make([]rune, 0, len(runes))
	n := 0

	x, i := f(runes[n:])
	ret = append(ret, x...)
	n += i
	for n < len(runes)-1 {
		x, i := f(runes[n:])
		ret = append(ret, x...)
		n += i
	}

	return string(append(ret, runes[n:]...))
}
