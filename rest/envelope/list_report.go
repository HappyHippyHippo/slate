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
func NewListReport(search string, start, count, total uint) *ListReport {
	prev := ""
	if start > 0 {
		nstart := uint(0)
		if count < start {
			nstart = start - count
		}
		prev = fmt.Sprintf("?search=%s&start=%d&count=%d", search, nstart, count)
	}

	next := ""
	if start+count < total {
		next = fmt.Sprintf("?search=%s&start=%d&count=%d", search, start+count, count)
	}

	return &ListReport{
		Search: search,
		Start:  start,
		Count:  count,
		Total:  total,
		Prev:   prev,
		Next:   next,
	}
}
