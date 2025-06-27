package property

const (
	MIN_RANGE  = -1
	HALF_RANGE = 1
	MAX_RANGE  = 1
)

type (
	Range             float64        // from -1 to +1
	GroupedProperties map[string]any // will be string or range
)

var (
	properties = make(map[int]GroupedProperties)
)

func get(group int, name string) any {
	if props := properties[group]; props != nil {
		return props[name]
	}
	// repeat search with parent or scope???
	return nil
}

func GetRange(group int, name string) Range {
	if prop := get(group, name); prop != nil {
		if r, ok := prop.(Range); ok {
			return r
		}
	}
	return HALF_RANGE
}

func GetBool(group int, name string) bool {
	if prop := get(group, name); prop != nil {
		if b, ok := prop.(bool); ok {
			return b
		}
	}
	return false // seems best to be false, e.g. not pressed...
}

func set(group int, name string, val any) {
	props := properties[group]
	if props == nil {
		props = make(GroupedProperties)
		properties[group] = props
	}
	props[name] = val
}

func SetRange(group int, name string, rang Range) {
	set(group, name, rang)
}

func SetBool(group int, name string, b bool) {
	set(group, name, b)
}
