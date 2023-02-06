package envelope

import (
	"fmt"
	"strconv"
	"strings"
)

// StatusError defines the structure to manipulate an error structure
// that hold the information of an execution error and be assigned to the
// response status error list.
type StatusError struct {
	Service  int    `json:"-" xml:"-"`
	Endpoint int    `json:"-" xml:"-"`
	Param    int    `json:"-" xml:"-"`
	Error    string `json:"-" xml:"-"`
	Code     string `json:"code" xml:"code"`
	Message  string `json:"message" xml:"message"`
}

// NewStatusError instantiates a new error instance.
func NewStatusError(
	e any,
	msg string,
) *StatusError {
	return (&StatusError{
		Error:   fmt.Sprintf("%v", e),
		Message: msg,
	}).compose()
}

// SetService assigns a service code value to the error.
func (e *StatusError) SetService(
	val int,
) *StatusError {
	e.Service = val
	return e.compose()
}

// SetEndpoint assigns an endpoint code value to the error.
func (e *StatusError) SetEndpoint(
	val int,
) *StatusError {
	e.Endpoint = val
	return e.compose()
}

// SetParam assigns a parameter code value to the error.
func (e *StatusError) SetParam(
	param int,
) *StatusError {
	e.Param = param
	return e.compose()
}

// SetError assigns a error code value to the error.
func (e *StatusError) SetError(
	err any,
) *StatusError {
	e.Error = fmt.Sprintf("%v", err)
	return e.compose()
}

// SetMessage assigns a message to the error.
func (e *StatusError) SetMessage(
	msg string,
) *StatusError {
	e.Message = msg
	return e
}

// GetCode retrieves the composed code of the error
func (e StatusError) GetCode() string {
	return e.Code
}

// GetMessage retrieves the message associated to the error
func (e StatusError) GetMessage() string {
	return e.Message
}

func (e *StatusError) compose() *StatusError {
	cb := strings.Builder{}
	// compose the service section of the code
	if e.Service != 0 {
		cb.WriteString(fmt.Sprintf("s:%d", e.Service))
	}
	// compose the endpoint section of the code
	if e.Endpoint != 0 {
		if cb.Len() != 0 {
			cb.WriteString(".")
		}
		cb.WriteString(fmt.Sprintf("e:%d", e.Endpoint))
	}
	// compose the param section of the code
	if e.Param != 0 {
		if cb.Len() != 0 {
			cb.WriteString(".")
		}
		cb.WriteString(fmt.Sprintf("p:%d", e.Param))
	}
	// compose the error section of the code
	if e.Error != "" {
		if cb.Len() != 0 {
			cb.WriteString(".")
		}

		if i, err := strconv.Atoi(e.Error); err != nil {
			cb.WriteString(e.Error)
		} else {
			cb.WriteString(fmt.Sprintf("c:%d", i))
		}
	}
	// assign the written code string to the error code structure parameter
	e.Code = cb.String()
	return e
}
