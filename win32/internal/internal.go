package internal

import (
	"unsafe"

	"github.com/mkch/gw/win32"
	"github.com/mkch/gw/win32/sysutil"
	"golang.org/x/sys/windows"
)

var lzUser32 = windows.NewLazyDLL("User32.dll")
var lzKernel32 = windows.NewLazyDLL("Kernel32.dll")
var lzGetClipboardData = lzUser32.NewProc("GetClipboardData")
var lzSetClipboardData = lzUser32.NewProc("SetClipboardData")
var lzOpenClipboard = lzUser32.NewProc("OpenClipboard")
var lzCloseClipboard = lzUser32.NewProc("CloseClipboard")
var lzGetDesktopWindow = lzUser32.NewProc("GetDesktopWindow")
var lzGlobalLock = lzKernel32.NewProc("GlobalLock")
var lzGlobalAlloc = lzKernel32.NewProc("GlobalAlloc")
var lzGlobalUnlock = lzKernel32.NewProc("GlobalUnlock")
var lzGlobalFree = lzKernel32.NewProc("GlobalFree")
var lzlstrlenW = lzKernel32.NewProc("lstrlenW")

func StrlenW(str *win32.WCHAR) int {
	return sysutil.As[int](lzlstrlenW.Call(uintptr(unsafe.Pointer(str))))
}

func GlobalFree(mem win32.HGLOBAL) error {
	return sysutil.MustZero[win32.HGLOBAL](lzGlobalFree.Call(uintptr(mem)))
}

func GlobalUnlock(mem win32.HGLOBAL) error {
	return sysutil.MustTrue(lzGlobalUnlock.Call(uintptr(mem)))
}

type GlobalAllocFlag win32.UINT

const GMEM_MOVEABLE GlobalAllocFlag = 0x0002

func GlobalAlloc(flags GlobalAllocFlag, n win32.SIZE_T) (win32.HGLOBAL, error) {
	return sysutil.MustNotZero[win32.HGLOBAL](lzGlobalAlloc.Call(uintptr(flags), uintptr(n)))
}

func GlobalLock(h win32.HGLOBAL) (win32.PVOID, error) {
	return sysutil.MustNotZero[win32.PVOID](lzGlobalLock.Call(uintptr(h)))
}

func GetDesktopWindow() win32.HWND {
	return sysutil.As[win32.HWND](lzGetDesktopWindow.Call())
}

func OpenClipboard(hwndOwner win32.HWND) error {
	sysutil.MustTrue(lzOpenClipboard.Call(uintptr(hwndOwner)))
	return nil
}

func CloseClipboard() error {
	return sysutil.MustTrue(lzCloseClipboard.Call())
}

type ClipboardFormat win32.UINT

const CF_UNICODETEXT ClipboardFormat = 13

func GetClipboardData(format ClipboardFormat) (win32.HGLOBAL, error) {
	return sysutil.MustNotZero[win32.HGLOBAL](lzGetClipboardData.Call(uintptr(format)))
}

func SetClipboardData(format ClipboardFormat, handle win32.HGLOBAL) error {
	return sysutil.MustTrue(lzSetClipboardData.Call(uintptr(format), uintptr(handle)))
}
