package judge

import (
	"open-judge/core"
)

type Judge interface {
	// 컴파일/런타임 오류 발생 시 에러 반환. result에선 정답 여부 반환
	Judge(problem core.Problem, code string) (result bool, err error)
}
