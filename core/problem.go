package core

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
