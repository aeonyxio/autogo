package window

import (
	"syscall"
	"unsafe"

	"github.com/aeonyxio/autogo/rectangle"

	"github.com/lxn/win"
)

type Window struct{
	id uintptr
}

func New(id uintptr) *Window {
	return &Window{
		id,
	}
}

func GetByName(name string) *Window {
	return New(uintptr(win.FindWindow(nil, syscall.StringToUTF16Ptr(name))))
}

func (w *Window) EnsureWindowActive() {
	hwnd := win.HWND(w.id)
	win.ShowWindow(hwnd, win.SW_RESTORE)
	win.SetForegroundWindow(hwnd)
	win.SetActiveWindow(hwnd)
}

func (w *Window) GetRegion() *rectangle.Rectangle {
	rect := win.RECT{}
	hwnd := win.HWND(w.id)
	win.GetWindowRect(hwnd, &rect)
	return rectangle.New(int(rect.Left),int(rect.Top),int(rect.Right - rect.Left),int(rect.Bottom - rect.Top))
}

// GetHWND get foreground window hwnd
func GetHWND() win.HWND {
	hwnd := win.GetForegroundWindow()
	return hwnd
}

// SendInput send n input event
func SendInput(nInputs uint32, pInputs unsafe.Pointer, cbSize int32) uint32 {
	return win.SendInput(nInputs, pInputs, cbSize)
}

// SendMsg send a message with hwnd
func SendMsg(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	return win.SendMessage(hwnd, msg, wParam, lParam)
}

// SetActiveWindow set window active with hwnd
func SetActiveWindow(hwnd win.HWND) win.HWND {
	return win.SetActiveWindow(hwnd)
}

// SetFocus set window focus with hwnd
func SetFocus(hwnd win.HWND) win.HWND {
	return win.SetFocus(hwnd)
}

// GetMian get the main display hwnd
func GetMain() win.HWND {
	return win.GetActiveWindow()
}

// GetMianId get the main display id
func GetMainId() int {
	return int(GetMain())
}
