package judge_test

import (
	"open-judge/core"
	"open-judge/judge"
	"open-judge/runner"
	"testing"
)

var (
	defaultLimit = core.Limit{
		TimeLimit:   1,
		MemoryLimit: 5 * 1000 * 1000, // 5MB
	}
	mockProblem = core.Problem{
		TestCases: []core.TestCase{
			{
				Input:  "",
				Answer: runner.MockAnswer,
			},
		},
		Limit: defaultLimit,
	}
)

func TestJudge(t *testing.T) {
	judge := judge.New(runner.MockRunner{})
	res, err := judge.Judge(mockProblem, runner.MockCode)
	if err != nil {
		t.Fatal("failed to judge!", err)
	}
	if !res {
		t.Fatal("answer is not correct!")
	}
}
