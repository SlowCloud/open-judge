package main

import (
	"open-judge/src/interfaces"
	"testing"
)

var runners = []interfaces.Runner{YaegiRunner{}, LocalRunner{}}

func TestRun(t *testing.T) {
	for _, runner := range runners {
		res := runner.Run(`package main; import "fmt"; func main() {fmt.Print("hello!")}`)
		if res.IsError() {
			t.Fatal("failed to run!", res.GetError().Error())
		}
	}
}

func TestRun_Fail(t *testing.T) {
	for _, runner := range runners {
		result := runner.Run(`randomword`)
		if !result.IsError() {
			t.Fatal("failed to fail code! wtf")
		}
	}
}
