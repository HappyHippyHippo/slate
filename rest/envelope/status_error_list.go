package envelope

import (
	"encoding/xml"
)

// StatusErrorList defines a type of data  that holds a list
// of error structures.
type StatusErrorList []*StatusError

// MarshalXML serialize the error list into a xml string
func (s StatusErrorList) MarshalXML(
	e *xml.Encoder,
	start xml.StartElement,
) error {
	// encode the list starting tag
	_ = e.EncodeToken(start)
	// iterate through all the stored error
	for _, v := range s {
		// create the iterated error starting tag name
		name := xml.Name{Space: "", Local: "error"}
		// encode the error instance tag with the code and message attributes
		_ = e.EncodeToken(xml.StartElement{
			Name: name,
			Attr: []xml.Attr{
				{Name: xml.Name{Local: "code"}, Value: v.Code},
				{Name: xml.Name{Local: "message"}, Value: v.Message},
			},
		})
		// encode the terminating error tag
		_ = e.EncodeToken(xml.EndElement{Name: name})
	}
	// encode the terminating list tag
	_ = e.EncodeToken(xml.EndElement{Name: start.Name})
	_ = e.Flush()
	return nil
}
