package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	runner := YaegiRunner{}
	result := runner.Run(`package main; import "fmt"; func main() {fmt.Print("hello!")}`)
	if result.IsError() {
		t.Fatal("failed to compile and run code!", result.GetError(), result.GetOutput())
	}
	if result.GetOutput() != "hello!" {
		t.Fatalf("output is not same! the output: %s", result.GetOutput())
	}
	t.Logf("success to run code! result: %s", result.GetOutput())
}

func TestRunFail(t *testing.T) {
	runner := YaegiRunner{}
	result := runner.Run(`randomword`)
	if !result.IsError() {
		t.Fatal("failed to fail code! wtf")
	}
	t.Log("success to fail code!")
}
