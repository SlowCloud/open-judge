package core

import "time"

type Problem struct {
	Description string
	TestCases   []TestCase
	Limit       Limit
}

type Limit struct {
	// second
	TimeLimit   int
	MemoryLimit uint64
}

type TestCase struct {
	Input  string
	Answer string
}

// returns "almost" no limit.
func NoLimit() Limit {
	return Limit{
		TimeLimit:   int(time.Hour),
		MemoryLimit: 1e18,
	}
}
