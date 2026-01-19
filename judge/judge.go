package judge

import (
	"open-judge/core"
	"open-judge/runner"
)

type Judge interface {
	setRunner(runner runner.Runner)
	Judge(problem Problem, code core.Code) (JudgeResult, JudgeError)
}

type Problem interface {
	Description() Description
	TestCases() []TestCase
}

type Description interface{ String() string }

type TestCase interface {
	Input() core.Input
	Answer() Answer
}
type Answer interface{ String() string }

type JudgeResult interface{ Result() []string }
type JudgeError interface{ Error() string }
