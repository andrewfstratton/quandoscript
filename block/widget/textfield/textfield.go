package textfield

type TextField struct {
	name     string
	default_ string
	empty    string
}

func New(name string) *TextField {
	return &TextField{name: name}
}

func (t *TextField) Html() (txt string) {
	// txt = `<input data-quando-name='name' type='text' value='initial' placeholder='empty'
	//  data-quando-encode='normal'/>
	txt = `<input data-quando-name='` + t.name + `' type='text'`
	if t.default_ != "" {
		txt = txt + " value='" + t.default_ + "'"
	}
	if t.empty != "" {
		txt = txt + " placeholder='" + t.empty + "'"
	}
	txt = txt + `/>`
	return
}

func (t *TextField) Default(s string) *TextField {
	t.default_ = s
	return t
}

func (t *TextField) Empty(s string) *TextField {
	t.empty = s
	return t
}
