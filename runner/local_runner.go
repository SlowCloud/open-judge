package runner

import (
	"bytes"
	"context"
	"open-judge/core"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v4/process"
)

type LocalRunner struct{}

// Run implements Runner.
func (l LocalRunner) Run(code string, limit core.Limit) (Result, error) {
	return l.RunWithInput(code, "", limit)
}

func (l LocalRunner) RunWithInput(code string, input string, limit core.Limit) (Result, error) {
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

	err = cmd.Start()
	if err != nil {
		return Result{}, err
	}

	go l.watchMemoryLimit(ctx, cancel, cmd, limit)

	err = cmd.Wait()
	if err != nil {
		return Result{}, err
	}

	return Result{Log: buf.String(), TimeTaken: time.Second, MemoryUsed: -1}, nil
}

func (l LocalRunner) watchMemoryLimit(ctx context.Context, cancel context.CancelFunc, cmd *exec.Cmd, limit core.Limit) {
	p, err := process.NewProcess(int32(cmd.Process.Pid))
	if err != nil {
		cancel()
	}

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			mem, err := p.MemoryInfo()
			if err != nil {
				cancel()
				return
			}
			if mem.RSS > uint64(limit.MemoryLimit) {
				cancel()
				return
			}
		}
	}
}

var _ Runner = LocalRunner{}
