package main

import (
	"bytes"
	"open-judge/src/interfaces"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type YaegiRunner struct {
}

var _ interfaces.Runner = YaegiRunner{}

func (judger YaegiRunner) Run(code interfaces.Code) interfaces.Result {

	var stdout bytes.Buffer
	i := interp.New(interp.Options{Stdout: &stdout})
	i.Use(stdlib.Symbols)

	_, err := i.Eval(string(code))
	if err != nil {
		return ResultDump{output: err.Error(), success: false, err: err}
	}

	return ResultDump{output: stdout.String(), success: true}
}
