package param

type Param struct {
	Val   any
	Qtype int
}

type Params map[string]Param
