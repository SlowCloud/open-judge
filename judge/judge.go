package judge

import (
	"open-judge/core"
	"open-judge/runner"
)

type Judge interface {
	setRunner(runner runner.Runner)

	// 컴파일/런타임 오류 발생 시 에러 반환. result에선 정답 여부 반환
	Judge(problem core.Problem, code core.Code) (result bool, err error)
}
