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

		for _, scenario := range scenarios {
			report := NewListReport(scenario.search, scenario.start, scenario.count, scenario.total)

			if check := report.Search; check != scenario.search {
				t.Errorf("stored the (%v) search terms when expecting (%v)", check, scenario.search)
			} else if check := report.Start; check != scenario.start {
				t.Errorf("stored the (%v) listing start record when expecting (%v)", check, scenario.start)
			} else if check := report.Count; check != scenario.count {
				t.Errorf("stored the (%v) listing record count when expecting (%v)", check, scenario.count)
			} else if check := report.Total; check != scenario.total {
				t.Errorf("stored the (%v) listing record total when expecting (%v)", check, scenario.total)
			} else if check := report.Prev; check != scenario.prev {
				t.Errorf("stored the (%v) prev link when expecting (%v)", check, scenario.prev)
			} else if check := report.Next; check != scenario.next {
				t.Errorf("stored the (%v) prev link when expecting (%v)", check, scenario.next)
			}
		}
	})
}
