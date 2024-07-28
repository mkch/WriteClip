package win32

import (
	"unsafe"

	"github.com/mkch/writeclip/win32/internal"

	"github.com/mkch/gw/win32"
	"github.com/mkch/gw/win32/win32util"
)

// ClipboardText gets the text in clipboard.
func ClipboardText() (str string, err error) {
	if err = internal.OpenClipboard(internal.GetDesktopWindow()); err != nil {
		return
	}
	defer internal.CloseClipboard()
	var mem win32.HGLOBAL
	if mem, err = internal.GetClipboardData(internal.CF_UNICODETEXT); err != nil {
		return
	}
	var ptr win32.PVOID
	if ptr, err = internal.GlobalLock(mem); err != nil {
		return
	}

	str = win32util.GoString((*win32.WCHAR)(ptr), internal.StrlenW((*win32.WCHAR)(ptr))+1)
	return
}

// SetClipboardText set str into clipboard.
func SetClipboardText(str string) (err error) {
	if err = internal.OpenClipboard(internal.GetDesktopWindow()); err != nil {
		return
	}
	defer internal.CloseClipboard()
	if err = internal.EmptyClipboard(); err != nil {
		return
	}
	var mem win32.HGLOBAL
	if mem, err = internal.GlobalAlloc(internal.GMEM_MOVEABLE, win32.SIZE_T((len(str)+1)*2)); err != nil {
		return
	}
	defer internal.GlobalFree(mem)

	func() {
		var p win32.PVOID
		if p, err = internal.GlobalLock(mem); err != nil {
			return
		}
		defer internal.GlobalUnlock(mem)
		dest := unsafe.Slice((*win32.WCHAR)(p), len(str)+1)
		win32util.CString(str, &dest)
	}()
	if err != nil {
		return
	}

	return internal.SetClipboardData(internal.CF_UNICODETEXT, mem)
}
