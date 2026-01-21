package core

import "time"

type Problem struct {
	Description string
	TestCases   []TestCase
	Limit       Limit
}

type Limit struct {
	// millisecond
	TimeLimit int64
	// byte
	MemoryLimit uint64
}

type TestCase struct {
	Input  string
	Answer string
}

// returns "almost" no limit.
func NoLimit() Limit {
	return Limit{
		TimeLimit:   int64(time.Hour),
		MemoryLimit: 1e18,
	}
}
