package judge

import (
	"open-judge/core"
	"open-judge/runner"
)

type Judge interface {
	setRunner(runner runner.Runner)
	Judge(problem core.Problem, code core.Code) (bool, error)
}
