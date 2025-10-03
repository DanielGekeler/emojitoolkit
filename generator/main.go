package main

import (
	"emoji-toolkit/internal"
	"encoding/binary"
	"encoding/xml"
	"os"
	"strconv"
)

func main() {
	repertoire := loadRepertoire("ucd.nounihan.flat.xml")
	GenerateEmojiRanges(repertoire)
	GenerateEmojiRanges2(repertoire)
	GenerateEmojiRanges3(repertoire)
}

// Singe Codepoint emojis with Emoji_Presentation=Yes & Emoji_Component=No
func GenerateEmojiRanges(repertoire internal.AnyXML) {
	codepoints := make([]int32, 0, 1024)

	for _, char := range repertoire.Children {
		if char.GetAttr("EPres") == "Y" && char.GetAttr("EComp") == "N" {
			n, _ := strconv.ParseUint(char.GetAttr("cp"), 16, 32)
			codepoints = append(codepoints, int32(n))
		}
	}

	writeRanges("emoji_ranges.bin", codepoints)
}

// Emojis that appear as text by default but can appear with an emoji presentation.
// These characters will appear as emojis when followed by U+FE0F (Variation Selector-16)
//
// ED-7 see https://www.unicode.org/reports/tr51/#def_text_presentation
func GenerateEmojiRanges2(repertoire internal.AnyXML) {
	codepoints := make([]int32, 0, 1024)

	for _, char := range repertoire.Children {
		if char.GetAttr("Emoji") == "Y" && char.GetAttr("EPres") == "N" {
			n, _ := strconv.ParseUint(char.GetAttr("cp"), 16, 32)
			codepoints = append(codepoints, int32(n))
		}
	}

	writeRanges("emoji_ranges2.bin", codepoints)
}

// Emojis that can be used in a RGI_Emoji_Modifier_Sequence
//
// ED-22 see https://www.unicode.org/reports/tr51/#def_std_emoji_modifier_sequence_set
func GenerateEmojiRanges3(repertoire internal.AnyXML) {
	codepoints := make([]int32, 0, 1024)

	for _, char := range repertoire.Children {
		if char.GetAttr("EBase") == "Y" {
			n, _ := strconv.ParseUint(char.GetAttr("cp"), 16, 32)
			if n != 0x1F46A {
				// Skip U+1F46A Family
				codepoints = append(codepoints, int32(n))
			}
		}
	}

	writeRanges("emoji_ranges3.bin", codepoints)
}

func loadRepertoire(path string) internal.AnyXML {
	data, _ := os.ReadFile(path)
	var ucd internal.AnyXML
	err := xml.Unmarshal(data, &ucd)
	if err != nil {
		panic("Error unmarshaling XML: " + err.Error())
	}

	var repertoire internal.AnyXML
	for _, v := range ucd.Children {
		if v.XMLName.Local == "repertoire" {
			repertoire = v
		}
	}
	return repertoire
}

func writeRanges(path string, codepoints []int32) {
	emoji_ranges := make([][]int32, 0)
	current_range := make([]int32, 0)
	for _, v := range codepoints {
		if len(current_range) == 0 {
			current_range = append(current_range, v)
			continue
		}

		if current_range[len(current_range)-1]+1 == v {
			current_range = append(current_range, v)
		} else {
			emoji_ranges = append(emoji_ranges, current_range)
			current_range = []int32{v}
		}
	}
	emoji_ranges = append(emoji_ranges, current_range)

	range_bytes := make([]byte, len(emoji_ranges)*8)
	for i, v := range emoji_ranges {
		binary.LittleEndian.PutUint32(range_bytes[i*8:i*8+4], uint32(v[0]))
		binary.LittleEndian.PutUint32(range_bytes[i*8+4:i*8+8], uint32(v[len(v)-1]))
	}
	os.WriteFile(path, range_bytes, os.ModePerm)
}
