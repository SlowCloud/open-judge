package core

type Problem struct {
	Description string
	TestCases   []TestCase
	TimeLimit   int
	MemoryLimit int
}

type TestCase struct {
	Input  string
	Answer string
}
