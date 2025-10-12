package internal

import (
	"encoding/xml"
	"os"
)

type AnyXML struct {
	XMLName  xml.Name
	Content  string     `xml:",chardata"`
	Attrs    []xml.Attr `xml:",any,attr"`
	Children []AnyXML   `xml:",any"`
}

func LoadXML(path string) AnyXML {
	data, _ := os.ReadFile(path)
	var anyXML AnyXML
	err := xml.Unmarshal(data, &anyXML)
	if err != nil {
		panic("Error unmarshaling XML: " + err.Error())
	}

	return anyXML
}

func (xml AnyXML) GetAttr(name string) string {
	for _, attr := range xml.Attrs {
		if attr.Name.Local == name {
			return attr.Value
		}
	}
	return ""
}

func (xml AnyXML) GetChildren(name string) []AnyXML {
	if name == "" {
		return xml.Children
	}

	ret := []AnyXML{}
	for _, child := range xml.Children {
		if child.XMLName.Local == name {
			ret = append(ret, child)
		}
	}
	return ret
}

func (xml AnyXML) GetFirstChild(name string) AnyXML {
	if name == "" && len(xml.Children) > 0 {
		return xml.Children[0]
	}

	for _, child := range xml.Children {
		if child.XMLName.Local == name {
			return child
		}
	}

	panic(name + " not found")
}
