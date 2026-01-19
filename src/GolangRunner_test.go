package main

import (
	"testing"
)

func TestJudge(t *testing.T) {
	runner := GolangRunner{}
	result := runner.Run(`print("testing judge!")`)
	if result.IsError() {
		t.Fatal("failed to compile and run code!")
	}
	if result.GetOutput() != "testing judge!" {
		t.Fatalf("output is not same! the output: %s", result.GetOutput())
	}
	t.Logf("success to run code! result: %s", result.GetOutput())
}

func TestJudgeFail(t *testing.T) {
	runner := GolangRunner{}
	result := runner.Run(`randomword`)
	if !result.IsError() {
		t.Fatal("failed to fail code! wtf")
	}
	t.Log("success to fail code!")
}
