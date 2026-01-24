## 설명

주어진 go 소스코드를 실행하여 메모리 한도, 제한시간 이내로 실행하는 라이브러리

## 요구사항

코드를 실행시키는 환경에 다음 실행 도구가 필요합니다.

- go
  - go 소스코드를 컴파일하고 실행시키기 위해 필요합니다.

## 타 패키지와의 차별점

- [yaegi](https://github.com/traefik/yaegi)
  - 소스코드 실행 옵션은 있으나, 실행시간/메모리 제한 기능은 없습니다.

## 패키지 구조

- core
  - 모든 패키지에서 주로 사용하는 인터페이스 / 구조체
- runner
  - 소스코드 실행 도구
  - runner 인터페이스를 주축으로 LocalRunner 구현체, Result 구조체 제공
- judge
  - 주어진 문제와 소스코드를 받아 실행하고 결과를 반환하는 구현체


## 패키지 활용법

### Runner

```golang
package main

import (
	"fmt"
	"open-judge/core"
	"open-judge/runner"
)

func main() {

	runner, err := runner.NewGo()
	if err != nil {
		panic(err)
	}

	code := `
package main

import "fmt"

func main() {
    fmt.Print("hello!")
}
`

	result, err := runner.Run(code, core.NoLimit())
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

```

### Judge

```golang
package main

import (
	"fmt"
	"open-judge/core"
	"open-judge/judge"
)

func main() {

	judge, err := judge.NewGo()
	if err != nil {
		panic(err)
	}

	code := `
package main

import "fmt"

func main() {
    fmt.Print("hello!")
}
`

	result, err := judge.Judge(core.Problem{TestCases: []core.TestCase{
		{
			Input:  "",
			Answer: "hello!",
		},
	},
		Limit: core.NoLimit(),
	}, code)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

```