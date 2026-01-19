package judge

import (
	"open-judge/core"
	"open-judge/runner"
)

type Judge interface {
	setRunner(runner runner.Runner)
	Judge(problem core.Problem, code core.Code) (JudgeResult, JudgeError)
}

type JudgeResult interface{ Result() []string }
type JudgeError interface{ Error() string }
