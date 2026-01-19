package main

import "testing"

func TestRun_LocalRunner(t *testing.T) {
	runner := LocalRunner{}
	res := runner.Run(`package main; import "fmt"; func main() {fmt.Print("hello!")}`)
	if res.IsError() {
		t.Fatal("failed to run!", res.GetError().Error())
	}
	t.Log(res.GetOutput())
}

func TestRun_Fail_LocalRunner(t *testing.T) {
	runner := LocalRunner{}
	result := runner.Run(`randomword`)
	if !result.IsError() {
		t.Fatal("failed to fail code! wtf")
	}
	t.Log("success to fail code!")
}
