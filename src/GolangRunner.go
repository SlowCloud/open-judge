package main

import (
	"bytes"
	"open-judge/src/interfaces"
	"strings"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type GolangRunner struct {
}

var _ interfaces.Runner = GolangRunner{}

func (judger GolangRunner) Run(code interfaces.Code) interfaces.Result {

	var stdout bytes.Buffer
	i := interp.New(interp.Options{Stdout: &stdout})
	i.Use(stdlib.Symbols)

	_, err := i.Eval(string(code))
	if err != nil {
		return DefaultResult{result: "", isError: true}
	}

	return DefaultResult{result: stdout.String(), isError: false}
}

type DefaultResult struct {
	result  string
	isError bool
}

var _ interfaces.Result = DefaultResult{}

func (r DefaultResult) GetOutput() string {
	return strings.Clone(r.result)
}

func (r DefaultResult) IsError() bool {
	return r.isError
}
