package main

import (
	"bytes"
	"open-judge/src/interfaces"
	"strings"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type YaegiRunner struct {
}

var _ interfaces.Runner = YaegiRunner{}

func (y YaegiRunner) Run(code interfaces.Code) (interfaces.Result, error) {
	return y.RunWithInput("", code)
}

func (y YaegiRunner) RunWithInput(input interfaces.Input, code interfaces.Code) (interfaces.Result, error) {
	stdin := strings.NewReader(string(input))
	var stdout bytes.Buffer
	i := interp.New(interp.Options{Stdin: stdin, Stdout: &stdout})
	i.Use(stdlib.Symbols)

	_, err := i.Eval(string(code))
	if err != nil {
		return nil, err
	}

	return ResultDump{output: stdout.String()}, nil
}
