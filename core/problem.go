package core

type Problem struct {
	Description string
	TestCases   []TestCase
	Limit       Limit
}

type Limit struct {
	TimeLimit   int
	MemoryLimit int
}

type TestCase struct {
	Input  string
	Answer string
}
