package emojitoolkit

import (
	_ "embed"
	"encoding/binary"
)

// Supported Unicode version
const Version = "16.0.0"

//go:generate go run generator/main.go

//go:embed emoji_ranges.bin
var emoji_ranges []byte

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
	return isInRange(r, emoji_ranges)
}

//go:embed emoji_ranges2.bin
var emoji_ranges2 []byte

// Matches default emoji presentation character ([ED-6]) as well as
// emoji presentation sequence ([ED-9a]). This should include all basic emoji defined by [ED-20].
//
// See section Basic_Emoji and Emoji_Keycap_Sequence of [emoji-sequences.txt] for a list of matching emojis.
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
// [ED-20]: https://www.unicode.org/reports/tr51/#def_basic_emoji_set
// [emoji-sequences.txt]: https://www.unicode.org/Public/emoji/latest/emoji-sequences.txt
func ContainsEmoji(s string) bool {
	runes := []rune(s)
	l := len(runes)
	for i, r := range runes {
		if IsSingleCharacterEmoji(r) {
			return true
		}

		if i+1 < l && isInRange(r, emoji_ranges2) && runes[i+1] == '\uFE0F' {
			return true
		}
	}
	return false
}

// Check if a rune is within any range of a set of ranges.
// These sets are an array of int32 with each range being defined
// as two consecutive int32 defining the first and last element
// of a range respectively.
func isInRange(r rune, ranges []byte) bool {
	for i := 0; i < len(ranges); i += 8 {
		min := binary.LittleEndian.Uint32(ranges[i:])
		max := binary.LittleEndian.Uint32(ranges[i+4:])
		if r >= rune(min) && r <= rune(max) {
			return true
		}
	}
	return false
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
	const flagA = 0x1F1E6 // REGIONAL INDICATOR SYMBOL LETTER A
	const flagB = 0x1F1FF // REGIONAL INDICATOR SYMBOL LETTER Z
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
