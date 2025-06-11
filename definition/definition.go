package definition

type Definition interface {
	TypeName() string
	Class() string
}

type (
	BOOLEAN  = bool
	STRING   = string
	NUMBER   = float64
	LINEID   = int
	UNKNOWN  = struct{}
	VARIABLE string // n.b. is not an alias which may cause extra code
	MENU     = string
	COLOUR   = string
	INTEGER  = int
	PERCENT  = float64
)
