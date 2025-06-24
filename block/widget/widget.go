package widget

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/andrewfstratton/quandoscript/block/widget/menuinput"
	"github.com/andrewfstratton/quandoscript/parse"
)

type (
	Widget interface {
		Html() string
	}
	Tag struct {
		Key string
		Val string
	}
)

func SetFields(widget any, tag string) {
	// using reflection to set fields
	v := reflect.ValueOf(widget).Elem() // i.e. pointer to struct

	tagList, err := tagToList(tag) // sorted list of struct tags in declaration order

	if err != nil {
		fmt.Println("error :", err)
		return
	}
	for _, tag := range tagList { // iterate through the tags
		key := tag.Key
		val := tag.Val
		vField := v.FieldByName(key)
		if !vField.CanSet() {
			if v.Type().Name() == "MenuInt" { // need to set map[int]string
				mi, ok := widget.(*menuinput.MenuInt)
				if ok {
					i, err := strconv.Atoi(key)
					if err == nil {
						mi.Choices = append(mi.Choices, menuinput.IntStr{Key: i, Val: val})
						continue
					}
				}
			}
			fmt.Printf("SetFields cannot set field '%s' in widget type '%s' with value '%s'\n", key, v.Type().Name(), val)
		} else {
			switch vField.Type().Name() {
			case "string":
				vField.SetString(val)
			case "bool":
				vField.SetBool(val == "true")
			case "Pfloat":
				f, err := strconv.ParseFloat(val, 64)
				if err != nil {
					fmt.Printf("Error parsing float for field '%s': %v\n", val, err)
					continue
				}
				vField.Set(reflect.ValueOf(&f))
			case "Pint":
				i, err := strconv.ParseInt(val, 10, 64)
				if err != nil {
					fmt.Printf("Error parsing int for field '%s': %v\n", key, err)
					continue
				}
				vField.Set(reflect.ValueOf(&i))
			default:
				fmt.Printf("Unknown type '%s' for field '%s' with value '%s'\n", vField.Type().Name(), key, val)
			}
		}
	}
}

func tagToList(tag string) (tagList []Tag, err error) {
	input := parse.Input{Line: tag}
	tagList = make([]Tag, 0)
	for input.Line != "" {
		key := input.GetTagKey() // ends when it runs out of letter/digit/_/. which will be at the ':' separator
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
		tag := Tag{Key: key, Val: val}
		tagList = append(tagList, tag)
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
