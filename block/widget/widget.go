package widget

type Widget interface {
	Html() string
}
func TagText(txt string, tag string) string {
	return OpenCloseTag(txt, tag, tag)
}

func OpenCloseTag(txt string, open string, close string) string {
	return fmt.Sprintf("<%v>%v</%v>", open, txt, close)
}
