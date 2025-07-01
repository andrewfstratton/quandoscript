package definition

import (
	"fmt"
	"reflect"
)

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

// Below is seperate since otherwise a copy is made of the definition and the copy only is changed
// unlike the block that is a new copy that is returned anyway...
func Setup(defn any) {
	typeDefn := reflect.TypeOf(defn).Elem()
	valueDefn := reflect.ValueOf(defn).Elem()
	for i := range typeDefn.NumField() {
		field := valueDefn.Field(i)
		nameTag := typeDefn.Field(i).Tag.Get("name")
		if nameTag == "" {
			continue
		}
		// use reflection to set name field
		if field.CanSet() {
			vName := field.FieldByName("Name")
			if vName.CanSet() {
				vName.SetString(nameTag)
				continue
			}
			fmt.Printf("Cannot set Name on %T\n", defn)
			continue
		}
		fmt.Printf("Cannot set field '%s' on %T\n", nameTag, defn)
	}
}
