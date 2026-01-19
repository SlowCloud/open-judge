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
