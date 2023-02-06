package envelope

import (
	"encoding/xml"
)

// Envelope identifies the structure of a response structured format.
type Envelope struct {
	XMLName    xml.Name    `json:"-" xml:"envelope"`
	StatusCode int         `json:"-" xml:"-"`
	Status     *Status     `json:"status" xml:"status"`
	ListReport *ListReport `json:"report,omitempty" xml:"report,omitempty"`
	Data       interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

// NewEnvelope instantiates a new response data envelope structure
func NewEnvelope(
	statusCode int,
	data interface{},
	listReport ...*ListReport,
) *Envelope {
	// initialize the envelope structure
	env := &Envelope{
		StatusCode: statusCode,
		Status:     NewStatus(),
		ListReport: nil,
		Data:       data,
	}
	// assign the list report if given as argument
	if len(listReport) > 0 && listReport[0] != nil {
		env.ListReport = listReport[0]
	}
	return env
}

// GetStatusCode returned the stored enveloped response status code
func (s Envelope) GetStatusCode() int {
	return s.StatusCode
}

// SetService assign the service identifier to all stored error codes
func (s *Envelope) SetService(
	val int,
) *Envelope {
	s.Status = s.Status.SetService(val)
	return s
}

// SetEndpoint assign the endpoint identifier to all stored error codes
func (s *Envelope) SetEndpoint(
	val int,
) *Envelope {
	s.Status = s.Status.SetEndpoint(val)
	return s
}

// SetListReport assign the list report to the envelope
func (s *Envelope) SetListReport(
	listReport *ListReport,
) *Envelope {
	s.ListReport = listReport
	return s
}

// AddError add a new error to the response envelope instance
func (s *Envelope) AddError(
	e *StatusError,
) *Envelope {
	s.Status = s.Status.AddError(e)
	return s
}
