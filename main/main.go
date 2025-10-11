package main

import (
	"fmt"

	"github.com/DanielGekeler/emojitoolkit"
)

// https://www.unicode.org/reports/tr42/
// https://www.unicode.org/reports/tr44/
// https://www.unicode.org/reports/tr51/#Identification
// https://www.unicode.org/glossary/
// https://www.unicode.org/Public/17.0.0/ucdxml/ucd.nounihan.flat.zip

func main() {
	fmt.Println(emojitoolkit.ContainsEmoji("1"))

	/*data, _ := os.ReadFile("../ucd.nounihan.flat.xml")
	var ucd internal.AnyXML
	err := xml.Unmarshal(data, &ucd)
	if err != nil {
		fmt.Println("Error unmarshaling XML:", err)
		return
	}

	var repertoire internal.AnyXML
	for _, v := range ucd.Children {
		if v.XMLName.Local == "repertoire" {
			repertoire = v
		}
	}

	for _, char := range repertoire.Children {
		if char.GetAttr("Emoji") == "Y" && char.GetAttr("EPres") == "N" {
			n, _ := strconv.ParseUint(char.GetAttr("cp"), 16, 32)

			fmt.Println(string(rune(n)))
			// fmt.Printf("%c %c\uFE0F\n", n, n)
		}
	}*/
}
