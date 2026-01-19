package core

type Problem interface {
	Description() string
	TestCases() []TestCase
}

type TestCase interface {
	Input() string
	Answer() string
}
