package widget

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/andrewfstratton/quandoscript/parse"
)

type Widget interface {
	Html() string
}

func Setup(widget any, name string, tag string) {
	// using reflection to set fields
	v := reflect.ValueOf(widget).Elem() // i.e. pointer to struct
	if name != "_" && name != "" {
		// using reflection to set name
		vName := v.FieldByName("Name")
		if vName.CanSet() {
			name = strings.ToLower(name[:1]) + name[1:] // lower case first letter
			vName.SetString(name)
		} else {
			fmt.Printf("Cannot set Name field on %T\n", widget)
		}
	}

	tagMap, err := tagToMap(tag)

	if err != nil {
		fmt.Println("error :", err)
		return
	}
	for key, str := range tagMap {
		vField := v.FieldByName(key)
		if vField.CanSet() {
			switch vField.Type().Name() {
			case "string":
				vField.SetString(str)
			case "bool":
				vField.SetBool(str == "true")
			case "Pfloat":
				f, err := strconv.ParseFloat(str, 64)
				if err != nil {
					fmt.Printf("Error parsing float for field '%s': %v\n", key, err)
					continue
				}
				vField.Set(reflect.ValueOf(&f))
			case "Pint":
				i, err := strconv.ParseInt(str, 10, 64)
				if err != nil {
					fmt.Printf("Error parsing int for field '%s': %v\n", key, err)
					continue
				}
				vField.Set(reflect.ValueOf(&i))
			default:
				fmt.Printf("Unknown type '%s' for field '%s' with value '%s'\n", vField.Type().Name(), key, str)
			}
		}
		// Note: TypeName and Class exist in Defn - not in widgets and txt used for widget txt to show
	}

}

func tagToMap(tag string) (tagMap map[string]string, err error) {
	input := parse.Input{Line: tag}
	tagMap = make(map[string]string)
	for input.Line != "" {
		key := input.GetWord() // ends when it runs out of letter/digit/_ which is by chance the same as :?!
		if input.Err != nil {
			err = input.Err
			return
		}
		key = strings.ToUpper(key[0:1]) + key[1:] // upper case first letter
		if input.GetColonDoublequote(); input.Err != nil {
			err = input.Err
			return
		}
		val := input.GetString()
		if input.Err != nil {
			err = input.Err
			return
		}
		tagMap[key] = val
		// this needs to be done so empty string detected correctly on next pass
		input.StripSpacer() // Note: ignores error if missing, i.e. at start of line
	}
	return
}

func TagText(txt string, tag string) string {
	return OpenCloseTag(txt, tag, tag)
}

func OpenCloseTag(txt string, open string, close string) string {
	return fmt.Sprintf("<%v>%v</%v>", open, txt, close)
}
