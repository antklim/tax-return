package taxreturn_test

import (
	"time"
)

type testPeriod struct {
	start time.Time
	end   time.Time
}

func (p testPeriod) Start() time.Time {
	return p.start
}

func (p testPeriod) End() time.Time {
	return p.end
}

func (p testPeriod) String() string {
	return ""
}

type periodCompareTestCase struct {
	desc          string
	p1            testPeriod
	p2            testPeriod
	within        bool
	outside       bool
	overlapsStart bool
	overlapsEnd   bool
}

var periodCompareTestCases = []periodCompareTestCase{
	{
		desc: "periods do not overlap",
		p1: testPeriod{
			start: time.Date(2020, time.January, 01, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.January, 02, 0, 0, 0, 0, time.UTC),
		},
		p2: testPeriod{
			start: time.Date(2020, time.February, 01, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.February, 02, 0, 0, 0, 0, time.UTC),
		},
		within:        false,
		outside:       true,
		overlapsStart: false,
		overlapsEnd:   false,
	},
	{
		desc: "period p1 within p2 (1)",
		p1: testPeriod{
			start: time.Date(2020, time.January, 02, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.January, 03, 0, 0, 0, 0, time.UTC),
		},
		p2: testPeriod{
			start: time.Date(2020, time.January, 01, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.January, 04, 0, 0, 0, 0, time.UTC),
		},
		within:        true,
		outside:       false,
		overlapsStart: false,
		overlapsEnd:   false,
	},
	{
		desc: "period p1 within p2 (2)",
		p1: testPeriod{
			start: time.Date(2020, time.January, 01, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.January, 04, 0, 0, 0, 0, time.UTC),
		},
		p2: testPeriod{
			start: time.Date(2020, time.January, 01, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.January, 04, 0, 0, 0, 0, time.UTC),
		},
		within:        true,
		outside:       false,
		overlapsStart: false,
		overlapsEnd:   false,
	},
	{
		desc: "period p1 within p2 (3)",
		p1: testPeriod{
			start: time.Date(2020, time.January, 01, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.January, 03, 0, 0, 0, 0, time.UTC),
		},
		p2: testPeriod{
			start: time.Date(2020, time.January, 01, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.January, 04, 0, 0, 0, 0, time.UTC),
		},
		within:        true,
		outside:       false,
		overlapsStart: false,
		overlapsEnd:   false,
	},
	{
		desc: "period p1 within p2 (4)",
		p1: testPeriod{
			start: time.Date(2020, time.January, 02, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.January, 04, 0, 0, 0, 0, time.UTC),
		},
		p2: testPeriod{
			start: time.Date(2020, time.January, 01, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.January, 04, 0, 0, 0, 0, time.UTC),
		},
		within:        true,
		outside:       false,
		overlapsStart: false,
		overlapsEnd:   false,
	},
}
