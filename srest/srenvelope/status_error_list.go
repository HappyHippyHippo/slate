package srenvelope

import (
	"encoding/xml"
)

// StatusErrorList defines a type of data  that holds a list of error structures.
type StatusErrorList []*StatusError

// MarshalXML serialize the error list into a xml string
func (s StatusErrorList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	_ = e.EncodeToken(start)
	for _, v := range s {
		name := xml.Name{Space: "", Local: "error"}
		_ = e.EncodeToken(xml.StartElement{
			Name: name,
			Attr: []xml.Attr{
				{Name: xml.Name{Local: "code"}, Value: v.Code},
				{Name: xml.Name{Local: "message"}, Value: v.Message},
			},
		})
		_ = e.EncodeToken(xml.EndElement{Name: name})
	}
	_ = e.EncodeToken(xml.EndElement{Name: start.Name})
	_ = e.Flush()

	return nil
}
