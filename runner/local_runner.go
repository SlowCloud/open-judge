package runner

import (
	"bytes"
	"context"
	"open-judge/core"
	"os"
	"os/exec"
	"strings"
	"time"
)

type localRunner struct{}

// Run implements Runner.
func (l localRunner) Run(code string, limit core.Limit) (Result, error) {
	return l.RunWithInput(code, "", limit)
}

func (l localRunner) RunWithInput(code string, input string, limit core.Limit) (Result, error) {
	file, err := os.CreateTemp("", "openjudge-*.go")
	if err != nil {
		return Result{}, err
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString(code)
	if err != nil {
		return Result{}, err
	}
	file.Close()

	background := context.Background()
	ctx, cancel := context.WithTimeout(background, time.Duration(limit.TimeLimit)*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", file.Name())
	cmd.Stdin = strings.NewReader(input)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf

	err = cmd.Run()
	if err != nil {
		return Result{}, err
	}

	return Result{Log: buf.String(), TimeTaken: time.Second, MemoryUsed: -1}, nil
}

var _ Runner = localRunner{}
