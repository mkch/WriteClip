package clipboard

import "github.com/mkch/writeclip/clipboard/internal/darwin"

// Text gets the text in clipboard.
func Text() (str string, err error) {
	return darwin.ClipboardText()
}

// Set sets str into clipboard.
func SetText(str string) (err error) {
	return darwin.SetClipboardText(str)
}
