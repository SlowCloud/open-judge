package main

import (
	"open-judge/src/interfaces"
	"os"
	"os/exec"
	"strings"
)

type LocalRunner struct {
}

func (l LocalRunner) Run(code interfaces.Code) (interfaces.Result, error) {
	return l.RunWithInput("", code)
}

func (l LocalRunner) RunWithInput(input interfaces.Input, code interfaces.Code) (interfaces.Result, error) {
	file, err := os.CreateTemp("", "openjudge-*.go")
	if err != nil {
		return nil, err
	}
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(code))
	if err != nil {
		return nil, err
	}
	file.Close()

	cmd := exec.Command("go", "run", file.Name())
	cmd.Stdin = strings.NewReader(string(input))

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return ResultDump{
		output: string(output),
	}, nil
}

var _ interfaces.Runner = LocalRunner{}
