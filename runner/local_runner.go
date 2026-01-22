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

func NewGo() Runner {
	return LocalRunner{}
}

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
	ctx, cancel := context.WithTimeout(background, time.Duration(limit.TimeLimit)*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", file.Name())

	cmd.Stdin = strings.NewReader(input)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf

	start := time.Now()

	err = cmd.Start()
	if err != nil {
		return Result{}, err
	}

	result := Result{}
	go l.watchMemoryLimit(ctx, cancel, cmd, limit, func(currentMemory uint64) {
		result.MemoryUsed = max(result.MemoryUsed, currentMemory)
	})

	err = cmd.Wait()
	if err != nil {
		return Result{}, err
	}

	end := time.Now()

	result.Log = buf.String()
	result.TimeTaken = end.Sub(start)

	return result, nil
}

// 메모리 감시.
// 메모리가 limit을 초과하면 캔슬 후 종료. 그 이유로 컨텍스트가 종료되어도 종료.
// 최대 메모리 사용량은 result에 기록됨. result에 바로 기록하기보다는 consumer func를 보내는 게 더 좋을 것 같다
func (l LocalRunner) watchMemoryLimit(ctx context.Context, cancel context.CancelFunc, cmd *exec.Cmd, limit core.Limit, consume func(uint64)) {
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
			consume(mem.RSS)
		}
	}
}

var _ Runner = LocalRunner{}
