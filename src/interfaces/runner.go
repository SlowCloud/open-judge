package interfaces

type Runner interface {
	Run(code Code) Result
}

type Result interface {
	IsError() bool
	GetOutput() string
	GetError() error
}
