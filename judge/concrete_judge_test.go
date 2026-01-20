package judge_test

import (
	"open-judge/core"
	"open-judge/judge"
	"open-judge/runner"
	"testing"
)

var (
	mockProblem = core.Problem{
		TestCases: []core.TestCase{
			{
				Input:  "",
				Answer: runner.MockAnswer,
			},
		},
		TimeLimit:   1,
		MemoryLimit: 5000 * 1000, // 5MB
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
