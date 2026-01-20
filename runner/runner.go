package runner

type Runner interface {
	Run(code string, timeout int) (Result, error)
	RunWithInput(input string, code string, timeout int) (Result, error)
}
