package runner

import "open-judge/core"

type Runner interface {
	Run(code string, limit core.Limit) (Result, error)
	RunWithInput(code string, input string, limit core.Limit) (Result, error)
}
