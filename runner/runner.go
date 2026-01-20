package runner

import "open-judge/core"

type Runner interface {
	Run(code core.Code, timeout int) (Result, error)
	RunWithInput(input core.Input, code core.Code, timeout int) (Result, error)
}
