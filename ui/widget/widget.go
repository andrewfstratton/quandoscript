package widget

type Style struct {
	Italic  bool
	Bold    bool
	Iconify bool
}

type Widget struct {
	Style Style
}

func (w *Widget) Italic() *Widget {
	w.Style.Italic = true
	return w
}

func (w *Widget) Bold() *Widget {
	w.Style.Bold = true
	return w
}

func (w *Widget) Iconify() *Widget {
	w.Style.Iconify = true
	return w
}
