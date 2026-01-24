package judge_test

import (
	"open-judge/core"
	"open-judge/judge"
	"open-judge/runner"
	"testing"
)

var (
	defaultProblem = core.Problem{
		TestCases: []core.TestCase{
			{
				Input:  "",
				Answer: "hello, go!",
			},
		},
		Limit: core.Limit{
			TimeLimit:   1000,
			MemoryLimit: 5 * 1000 * 1000, // 5MB
		},
	}
	helloWorldCode = `
	package main
	import "fmt"
	func main() {fmt.Print("hello, go!")}
	`
)

func TestJudge(t *testing.T) {
	r, err := runner.NewGo()
	if err != nil {
		t.Fatal(err)
	}
	judge := judge.New(r)
	res, err := judge.Judge(defaultProblem, helloWorldCode)
	if err != nil {
		t.Fatal("failed to judge!", err)
	}
	if !res {
		t.Fatal("answer is not correct!")
	}
}
