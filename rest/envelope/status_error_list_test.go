package envelope

import (
	"encoding/xml"
	"strings"
	"testing"
)

func Test_StatusErrorList_MarshalXML(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		name := "start"
		buffer := strings.Builder{}
		start := xml.StartElement{Name: xml.Name{Local: name}}
		list := StatusErrorList{}
		expected := "<start></start>"

		if err := list.MarshalXML(xml.NewEncoder(&buffer), start); err != nil {
			t.Errorf("returned the ujnexpected error (%v)", err)
		} else if check := buffer.String(); check != expected {
			t.Errorf("marshaled the list into (%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("single element list", func(t *testing.T) {
		name := "start"
		buffer := strings.Builder{}
		start := xml.StartElement{Name: xml.Name{Local: name}}
		list := StatusErrorList{NewStatusError(1, "error message").SetService(2).SetEndpoint(3)}
		expected := `<start><error code="s:2.e:3.c:1" message="error message"></error></start>`

		if err := list.MarshalXML(xml.NewEncoder(&buffer), start); err != nil {
			t.Errorf("returned the ujnexpected error (%v)", err)
		} else if check := buffer.String(); check != expected {
			t.Errorf("marshaled the list into (%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("multiple element list", func(t *testing.T) {
		name := "start"
		buffer := strings.Builder{}
		start := xml.StartElement{Name: xml.Name{Local: name}}
		list := StatusErrorList{
			NewStatusError(1, "error message 1").SetService(2).SetEndpoint(3),
			NewStatusError(2, "error message 2").SetService(2).SetEndpoint(3),
			NewStatusError(3, "error message 3").SetService(2).SetEndpoint(3),
		}
		expected := `<start>`
		expected += `<error code="s:2.e:3.c:1" message="error message 1"></error>`
		expected += `<error code="s:2.e:3.c:2" message="error message 2"></error>`
		expected += `<error code="s:2.e:3.c:3" message="error message 3"></error>`
		expected += `</start>`

		if err := list.MarshalXML(xml.NewEncoder(&buffer), start); err != nil {
			t.Errorf("returned the ujnexpected error (%v)", err)
		} else if check := buffer.String(); check != expected {
			t.Errorf("marshaled the list into (%v) when expecting (%v)", check, expected)
		}
	})
}
