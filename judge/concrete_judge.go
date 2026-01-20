package judge

import (
	"errors"
	"open-judge/core"
	"open-judge/runner"
	"time"
)

type concreteJudge struct {
	runner runner.Runner
}

func New(runner runner.Runner) Judge {
	return concreteJudge{
		runner: runner,
	}
}

func (c concreteJudge) Judge(problem core.Problem, code string) (result bool, err error) {

	for _, testcase := range problem.TestCases {

		res, err := c.runner.RunWithInput(code, testcase.Input, problem.TimeLimit)
		if err != nil {
			return false, err
		}

		if res.MemoryUsed > problem.MemoryLimit {
			return false, errors.New("Memory Limit Exceed")
		}

		if res.TimeTaken > time.Duration(problem.TimeLimit)*time.Second {
			return false, errors.New("Time Limit Exceed")
		}

		if res.Log != testcase.Answer {
			return false, nil
		}
	}

	return true, nil
}

var _ Judge = concreteJudge{}
