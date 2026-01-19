package runner

import (
	"bytes"
	"open-judge/core"
	"strings"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type YaegiRunner struct {
}

var _ Runner = YaegiRunner{}

func (y YaegiRunner) Run(code core.Code) (Result, error) {
	return y.RunWithInput(core.EmptyInput(), code)
}

func (y YaegiRunner) RunWithInput(input core.Input, code core.Code) (Result, error) {
	stdin := strings.NewReader(input.String())
	var stdout bytes.Buffer
	i := interp.New(interp.Options{Stdin: stdin, Stdout: &stdout})
	i.Use(stdlib.Symbols)

	_, err := i.Eval(code.String())
	if err != nil {
		return nil, err
	}

	return concreteResult{output: stdout.String()}, nil
}
