package clipboard

import (
	"github.com/mkch/writeclip/clipboard/internal/win32"
)

// Text gets the text in clipboard.
func Text() (str string, err error) {
	return win32.ClipboardText()
}

// Set sets str into clipboard.
func SetText(str string) (err error) {
	return win32.SetClipboardText(str)
}
