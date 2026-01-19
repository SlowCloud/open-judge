package runner

import (
	"open-judge/core"
	"testing"
)

var runners = []Runner{YaegiRunner{}, LocalRunner{}}

func TestRun(t *testing.T) {
	for _, runner := range runners {
		_, err := runner.Run(core.NewCode(`package main; import "fmt"; func main() {fmt.Print("hello!")}`))
		if err != nil {
			t.Fatal("failed to run!", err)
		}
	}
}

func TestRun_Fail(t *testing.T) {
	for _, runner := range runners {
		_, err := runner.Run(core.NewCode(`randomword`))
		if err == nil {
			t.Fatal("failed to fail code!", err)
		}
	}
}

func TestRunWithInput(t *testing.T) {
	for _, runner := range runners {
		input := "this_is_input!"
		res, err := runner.RunWithInput(core.NewInput(input), core.NewCode(`package main; import "fmt"; func main() {
		var s string
		fmt.Scan(&s)
		fmt.Print(s)
		}`))
		if err != nil {
			t.Fatal("failed to run code!", err)
		}
		if res.GetOutput() != input {
			t.Fatal("output is not same with expected! ", res.GetOutput())
		}
	}
}
