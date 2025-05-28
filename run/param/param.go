package param

type Param struct {
	Val   any
	Qtype int
}

const (
	UNKNOWN int = iota
	VARIABLE
	BOOLEAN
	STRING
	NUMBER // may need range and integer
	LINEID
)

type Params map[string]Param
type Op func(Params) func(Params)
