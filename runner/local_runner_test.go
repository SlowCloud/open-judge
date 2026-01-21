package runner_test

import (
	"open-judge/core"
	"open-judge/runner"
	"testing"
)

func Test_LocalRunner_Run(t *testing.T) {
	runner_ := runner.LocalRunner{}
	testRun(runner_, t)
}

func Test_LocalRunner_Run_Timeout(t *testing.T) {
	runner_ := runner.LocalRunner{}
	_, err := runner_.Run(`package main; import "time"; func main() {time.Sleep(10*time.Minute)}`, core.Limit{TimeLimit: 1, MemoryLimit: -1})
	if err == nil {
		t.Fatal("it must fail...")
	}
	t.Log(err)
}

func Test_LocalRunner_Run_MemoryLimit(t *testing.T) {
	runner_ := runner.LocalRunner{}

	// 100MB 크기의 슬라이스를 할당하여 메모리 한도를 즉시 넘기도록 함
	targetCode := ` package main; import "time";
func main() {
    _ = make([]byte, 100 * 1024 * 1024) // 100MB 할당
    time.Sleep(2 * time.Second)          // 모니터링 루프가 감지할 시간 확보
}
`

	limit := core.Limit{
		TimeLimit:   5,
		MemoryLimit: 10 * 1024 * 1024, // 10MB 제한
	}

	_, err := runner_.Run(targetCode, limit)

	if err == nil {
		t.Fatal("메모리 초과로 인해 실패해야 하지만 에러가 발생하지 않았습니다.")
	}

	t.Logf("예상된 에러 발생: %v", err)
}

func Test_LocaRunner_Run_Fail(t *testing.T) {
	runner_ := runner.LocalRunner{}
	testRun_fail(runner_, t)
}
