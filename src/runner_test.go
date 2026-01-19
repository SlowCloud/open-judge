package main

import (
	"open-judge/src/interfaces"
	"testing"
)

var runners = []interfaces.Runner{YaegiRunner{}, LocalRunner{}}

func TestRun(t *testing.T) {
	for _, runner := range runners {
		_, err := runner.Run(`package main; import "fmt"; func main() {fmt.Print("hello!")}`)
		if err != nil {
			t.Fatal("failed to run!", err)
		}
	}
}

func TestRun_Fail(t *testing.T) {
	for _, runner := range runners {
		_, err := runner.Run(`randomword`)
		if err == nil {
			t.Fatal("failed to fail code!", err)
		}
	}
}

func TestRunWithInput(t *testing.T) {
	for _, runner := range runners {
		input := "this_is_input!"
		res, err := runner.RunWithInput(interfaces.Input(input), `package main; import "fmt"; func main() {
		var s string
		fmt.Scan(&s)
		fmt.Print(s)
		}`)
		if err != nil {
			t.Fatal("failed to run code!", err)
		}
		if res.GetOutput() != input {
			t.Fatal("output is not same with expected! ", res.GetOutput())
		}
	}
}
