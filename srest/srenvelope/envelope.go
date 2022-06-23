package srenvelope

import (
	"encoding/xml"
)

// Envelope identifies the structure of a response structured format.
type Envelope struct {
	XMLName    xml.Name    `json:"-" xml:"envelope"`
	StatusCode int         `json:"-" xml:"-"`
	Status     *Status     `json:"status" xml:"status"`
	Report     *ListReport `json:"report,omitempty" xml:"report,omitempty"`
	Data       interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

// NewEnvelope instantiates a new response data envelope structure
func NewEnvelope(statusCode int, data interface{}, report *ListReport) *Envelope {
	return &Envelope{
		StatusCode: statusCode,
		Status:     NewStatus(),
		Report:     report,
		Data:       data,
	}
}

// GetStatusCode returned the stored enveloped response status code
func (s Envelope) GetStatusCode() int {
	return s.StatusCode
}

// SetService assign the service identifier to all stored error codes
func (s *Envelope) SetService(val int) *Envelope {
	s.Status = s.Status.SetService(val)
	return s
}

// SetEndpoint assign the endpoint identifier to all stored error codes
func (s *Envelope) SetEndpoint(val int) *Envelope {
	s.Status = s.Status.SetEndpoint(val)
	return s
}

// AddError add a new error to the response envelope instance
func (s *Envelope) AddError(err *StatusError) *Envelope {
	s.Status = s.Status.AddError(err)
	return s
}
