package darwin

import "testing"

func TestSetClipboardText(t *testing.T) {
	if err := SetClipboardText(""); err != nil {
		t.Fatal(err)
	}
	if str, err := ClipboardText(); err != nil {
		t.Fatal(err)
	} else if str != "" {
		t.Fatal(err)
	}
	if err := SetClipboardText("abc"); err != nil {
		t.Fatal(err)
	}
	if str, err := ClipboardText(); err != nil {
		t.Fatal(err)
	} else if str != "abc" {
		t.Fatal(err)
	}
}
