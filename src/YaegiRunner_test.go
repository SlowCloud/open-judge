package main

import (
	"testing"
)

func TestRun_YaegiRunner(t *testing.T) {
	runner := YaegiRunner{}
	res := runner.Run(`package main; import "fmt"; func main() {fmt.Print("hello!")}`)
	if res.IsError() {
		t.Fatal("failed to run!", res.GetError().Error())
	}
	t.Log(res.GetOutput())
}

func TestRun_Fail_YaegiRunner(t *testing.T) {
	runner := YaegiRunner{}
	result := runner.Run(`randomword`)
	if !result.IsError() {
		t.Fatal("failed to fail code! wtf")
	}
	t.Log("success to fail code!")
}
