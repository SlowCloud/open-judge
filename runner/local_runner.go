package runner

import (
	"open-judge/core"
	"os"
	"os/exec"
	"strings"
)

type LocalRunner struct {
}

func (l LocalRunner) Run(code core.Code) (Result, error) {
	return l.RunWithInput(core.EmptyInput(), code)
}

func (l LocalRunner) RunWithInput(input core.Input, code core.Code) (Result, error) {
	file, err := os.CreateTemp("", "openjudge-*.go")
	if err != nil {
		return nil, err
	}
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(code.String()))
	if err != nil {
		return nil, err
	}
	file.Close()

	cmd := exec.Command("go", "run", file.Name())
	cmd.Stdin = strings.NewReader(input.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return concreteResult{
		output: string(output),
	}, nil
}

var _ Runner = LocalRunner{}
