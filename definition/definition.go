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
		vField := valueDefn.Field(i)
		name := typeDefn.Field(i).Name
		fmt.Printf("field name is '%s'\n", name)
		if name == "" || name == "_" {
			continue
		}
		// use reflection to set name field
		if vField.CanSet() {
			vName := vField.FieldByName("Name")
			fmt.Printf("...can set field, Vname '%s'\n", vName)
			if !vName.IsValid() { // i.e. not found/no value so skip
				continue
			}
			if vName.CanSet() {
				vName.SetString(name)
				fmt.Printf("...has set Vname to '%s'\n", name)
				continue
			}
			fmt.Printf("Cannot set Name on %s\n", name)
			continue
		}
		fmt.Printf("Cannot set field '%s' on %T\n", name, defn)
	}
}
