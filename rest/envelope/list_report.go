package envelope

import (
	"fmt"
)

// ListReport defines the structure of a response list report
// containing all the request information, but also the total amount of
// filtering records and links for the previous and next pages
type ListReport struct {
	Search string `json:"search" xml:"search"`
	Start  uint   `json:"start" xml:"start"`
	Count  uint   `json:"count" xml:"count"`
	Total  uint   `json:"total" xml:"total"`
	Prev   string `json:"prev" xml:"prev"`
	Next   string `json:"next" xml:"next"`
}

// NewListReport instantiates a new response list report by
// populating the prev and next link information regarding the given
// filtering information
func NewListReport(
	search string,
	start,
	count,
	total uint,
) *ListReport {
	// store the prev URL query parameters if the start value
	// is greater than zero
	prev := ""
	if start > 0 {
		// discover the previous page starting value
		nstart := uint(0)
		if count < start {
			nstart = start - count
		}
		// compose the URL prev page query parameters
		prev = fmt.Sprintf("?search=%s&start=%d&count=%d", search, nstart, count)
	}
	// store the next URL query parameters if the total number of
	// record are greater than the current start plus the number of
	// presented records
	next := ""
	if start+count < total {
		// compose the URL next page query parameters
		next = fmt.Sprintf("?search=%s&start=%d&count=%d", search, start+count, count)
	}
	// return the list report instance reference
	return &ListReport{
		Search: search,
		Start:  start,
		Count:  count,
		Total:  total,
		Prev:   prev,
		Next:   next,
	}
}
