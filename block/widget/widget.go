package widget

type Style struct {
	Italic  bool
	Bold    bool
	Iconify bool
}

type Widget interface {
	LinkTo(widget Widget)
}

type TextWidget struct {
	Style Style
	next  *Widget
}

func (t *TextWidget) Italic() *TextWidget {
	t.Style.Italic = true
	return t
}

func (t *TextWidget) Bold() *TextWidget {
	t.Style.Bold = true
	return t
}

func (t *TextWidget) Iconify() *TextWidget {
	t.Style.Iconify = true
	return t
}

func (t *TextWidget) LinkTo(widget Widget) {
	t.next = &widget
}
