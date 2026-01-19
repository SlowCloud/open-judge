package main

import (
	"open-judge/src/interfaces"
	"os"
	"os/exec"
	"strings"
)

type LocalRunner struct {
}

func (l LocalRunner) Run(code interfaces.Code) interfaces.Result {
	return l.RunWithInput("", code)
}

func (l LocalRunner) RunWithInput(input interfaces.Input, code interfaces.Code) interfaces.Result {
	file, err := os.CreateTemp("", "openjudge-*.go")
	if err != nil {
		return ResultDump{success: false, err: err}
	}
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(code))
	if err != nil {
		return ResultDump{success: false, err: err}
	}
	file.Close()

	cmd := exec.Command("go", "run", file.Name())
	cmd.Stdin = strings.NewReader(string(input))

	output, err := cmd.CombinedOutput()
	if err != nil {
		return ResultDump{
			success: false,
			output:  string(output),
			err:     err,
		}
	}

	return ResultDump{
		output:  string(output),
		success: true,
	}
}

var _ interfaces.Runner = LocalRunner{}
