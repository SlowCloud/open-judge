package core

type Problem interface {
	Description() Description
	TestCases() []TestCase
}

type Description interface{ String() string }

type TestCase interface {
	Input() TestInput
	Answer() TestAnswer
}

type TestInput interface{ String() string }
type TestAnswer interface{ String() string }
