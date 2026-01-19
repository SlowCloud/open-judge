package interfaces

type Judger interface {
	setRunner(runner Runner)
	Judge(problem Problem, code Code) JudgeResult
}

type Problem interface {
	Description() Description
	TestCases() []TestCase
}

type Description string

type TestCase interface {
	Input() Input
	Answer() Answer
}

type Answer string

type JudgeResult interface {
	Success() bool
	Fail() bool
	ErrorMessage() ErrorMessage
}

type ErrorMessage string
