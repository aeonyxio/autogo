package mouse

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/aeonyxio/autogo/point"

	"github.com/lxn/win"
)

type mouseButton int

const (
    LEFT mouseButton = iota
    RIGHT
    MIDDLE
)

type Mouse struct {
}

func New() *Mouse {
	return &Mouse{}
}

func (mouse *Mouse) Move(p *point.Point){
	win.SetCursorPos(int32(p.X), int32(p.Y))
}

func (mouse *Mouse) GetPosition() *point.Point {
	p := &win.POINT{}
	win.GetCursorPos(p)
	return point.New(int(p.X), int(p.Y))
}

func (mouse *Mouse) RightPress(){
	mouseDown(RIGHT)
}

func (mouse *Mouse) RightRelease(){
	mouseUp(RIGHT)
}

func (mouse *Mouse) RightClick(){
	mouse.RightPress()
    time.Sleep(20 * time.Millisecond)
	mouse.RightRelease()
}

func (mouse *Mouse) MiddlePress(){
	mouseDown(MIDDLE)
}

func (mouse *Mouse) MiddleRelease(){
	mouseUp(MIDDLE)
}

func (mouse *Mouse) MiddleClick(){
	mouse.MiddlePress()
    time.Sleep(20 * time.Millisecond)
	mouse.MiddleRelease()
}

func (mouse *Mouse) Press(){
	mouseDown(LEFT)
}

func (mouse *Mouse) Release(){
	mouseUp(LEFT)
}

func (mouse *Mouse) Click(){
	mouse.Press()
    time.Sleep(20 * time.Millisecond)
	mouse.Release()
}

func mouseDown(mouse mouseButton) {  
	in := []win.MOUSE_INPUT{
		{
			Type: win.INPUT_MOUSE,
			Mi: win.MOUSEINPUT{},
		},
	}

  	switch mouse {
    case LEFT:
        in[0].Mi.DwFlags = win.MOUSEEVENTF_LEFTDOWN
    case RIGHT:
        in[0].Mi.DwFlags = win.MOUSEEVENTF_RIGHTDOWN
    case MIDDLE:
        in[0].Mi.DwFlags = win.MOUSEEVENTF_MIDDLEDOWN
    default:
        fmt.Println("Woops")
    }

    win.SendInput(1,  unsafe.Pointer(&in[0]), int32(unsafe.Sizeof(in[0])));
}

func mouseUp(mouse mouseButton) {  
	in := []win.MOUSE_INPUT{
		{
			Type: win.INPUT_MOUSE,
			Mi: win.MOUSEINPUT{},
		},
	}

  	switch mouse {
		case LEFT:
			in[0].Mi.DwFlags = win.MOUSEEVENTF_LEFTUP
		case RIGHT:
			in[0].Mi.DwFlags = win.MOUSEEVENTF_RIGHTUP
		case MIDDLE:
			in[0].Mi.DwFlags = win.MOUSEEVENTF_MIDDLEUP
		default:
			fmt.Println("Woops")
    }

    win.SendInput(1,  unsafe.Pointer(&in[0]), int32(unsafe.Sizeof(in[0])));
}