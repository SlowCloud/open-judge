package interfaces

type Runner interface {
	Run(code Code) (Result, error)
}

type Result interface {
	GetOutput() string
}
