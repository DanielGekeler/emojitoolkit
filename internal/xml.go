package internal

import "encoding/xml"

type AnyXML struct {
	XMLName  xml.Name
	Content  string     `xml:",chardata"`
	Attrs    []xml.Attr `xml:",any,attr"`
	Children []AnyXML   `xml:",any"`
}

func (xml AnyXML) GetAttr(name string) string {
	for _, attr := range xml.Attrs {
		if attr.Name.Local == name {
			return attr.Value
		}
	}
	return ""
}
