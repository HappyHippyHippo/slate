package envelope

import "testing"

func Test_NewListReport(t *testing.T) {
	t.Run("store the search parameters", func(t *testing.T) {
		scenarios := []struct {
			search string
			start  uint
			count  uint
			total  uint
			prev   string
			next   string
		}{
			{ // report on start position
				search: "search string",
				start:  uint(0),
				count:  uint(2),
				total:  uint(10),
				prev:   "",
				next:   "?search=search string&start=2&count=2",
			},
			{ // report with truncated prev link
				search: "search string",
				start:  uint(1),
				count:  uint(2),
				total:  uint(10),
				prev:   "?search=search string&start=0&count=2",
				next:   "?search=search string&start=3&count=2",
			},
			{ // report with prev link
				search: "search string",
				start:  uint(2),
				count:  uint(2),
				total:  uint(10),
				prev:   "?search=search string&start=0&count=2",
				next:   "?search=search string&start=4&count=2",
			},
			{ // report with prev link (2)
				search: "search string",
				start:  uint(3),
				count:  uint(2),
				total:  uint(10),
				prev:   "?search=search string&start=1&count=2",
				next:   "?search=search string&start=5&count=2",
			},
			{ // report without next page
				search: "search string",
				start:  uint(8),
				count:  uint(2),
				total:  uint(10),
				prev:   "?search=search string&start=6&count=2",
				next:   "",
			},
			{ // report without next page (2)
				search: "search string",
				start:  uint(9),
				count:  uint(2),
				total:  uint(10),
				prev:   "?search=search string&start=7&count=2",
				next:   "",
			},
			{ // report without next page (3)
				search: "search string",
				start:  uint(10),
				count:  uint(2),
				total:  uint(10),
				prev:   "?search=search string&start=8&count=2",
				next:   "",
			},
		}

		for _, s := range scenarios {
			report := NewListReport(s.search, s.start, s.count, s.total)

			if check := report.Search; check != s.search {
				t.Errorf("(%v) when expecting (%v)", check, s.search)
			} else if check := report.Start; check != s.start {
				t.Errorf("(%v) when expecting (%v)", check, s.start)
			} else if check := report.Count; check != s.count {
				t.Errorf("(%v) when expecting (%v)", check, s.count)
			} else if check := report.Total; check != s.total {
				t.Errorf("(%v) when expecting (%v)", check, s.total)
			} else if check := report.Prev; check != s.prev {
				t.Errorf("(%v) when expecting (%v)", check, s.prev)
			} else if check := report.Next; check != s.next {
				t.Errorf("(%v) when expecting (%v)", check, s.next)
			}
		}
	})
}
