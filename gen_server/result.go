package gen_server

type ResultType int

const (
	ReasonNormal ResultType = iota
	ReasonShutdown
	OptNormal
	OptTerminate
	ResultOk
	ResultError
)

type Result interface {
	Type() ResultType
	Detail() interface{}
}

func (r ResultType) Type() ResultType {
	return ResultType(r)
}

func (r ResultType) Detail() interface{} {
	return r
}
