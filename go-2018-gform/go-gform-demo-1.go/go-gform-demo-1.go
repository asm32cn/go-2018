// go-gform-demo-1.go
package main

import (
	"fmt"
	"unsafe"
	"math/rand"
	"github.com/AllenDang/gform"
	// "github.com/AllenDang/w32"
	"syscall"
)

var (
	modgdi32            = syscall.NewLazyDLL("gui32.dll")
	moduser32           = syscall.NewLazyDLL("user32.dll")

	procGetDC           = moduser32.NewProc("GetDC")
	procReleaseDC       = moduser32.NewProc("ReleaseDC")
	procInvalidateRect  = moduser32.NewProc("InvalidateRect")
	procSetTimer        = moduser32.NewProc("SetTimer")
	procGetClientRect   = moduser32.NewProc("GetClientRect")
	procPostQuitMessage = moduser32.NewProc("PostQuitMessage")

	procLineTo          = modgdi32.NewProc("LineTo")
	procMoveToEx        = modgdi32.NewProc("MoveToEx")
)

const (
	WM_CLOSE                  = 16
	WM_SIZE                   = 5
	WM_TIMER                  = 275
)

type POINT struct {
	X, Y int32
}

type RECT struct {
	Left, Top, Right, Bottom int32
}

type (
	BOOL            int32
	HANDLE          uintptr
	HWND            HANDLE
	HDC             HANDLE
	COLORREF        uint32
)

func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}

	return 0
}

func PostQuitMessage(exitCode int) {
	procPostQuitMessage.Call(
		uintptr(exitCode))
}

func GetDC(hwnd HWND) HDC {
	ret, _, _ := procGetDC.Call(
		uintptr(hwnd))

	return HDC(ret)
}

func ReleaseDC(hwnd HWND, hDC HDC) bool {
	ret, _, _ := procReleaseDC.Call(
		uintptr(hwnd),
		uintptr(hDC))

	return ret != 0
}

func SetTimer(hwnd HWND, nIDEvent uint32, uElapse uint32, lpTimerProc uintptr) uintptr {
	ret, _, _ := procSetTimer.Call(
		uintptr(hwnd),
		uintptr(nIDEvent),
		uintptr(uElapse),
		lpTimerProc,
	)
	return ret
}

func GetClientRect(hwnd HWND) *RECT {
	var rect RECT
	ret, _, _ := procGetClientRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect)))

	if ret == 0 {
		panic(fmt.Sprintf("GetClientRect(%d) failed", hwnd))
	}

	return &rect
}

func LineTo(hdc HDC, nXEnd, nYEnd int) bool {
	ret, _, _ := procLineTo.Call(
		uintptr(hdc),
		uintptr(nXEnd),
		uintptr(nYEnd))

	return ret != 0
}

func InvalidateRect(hwnd HWND, rect *RECT, erase bool) bool {
	ret, _, _ := procInvalidateRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(rect)),
		uintptr(BoolToBOOL(erase)))

	return ret != 0
}

func MoveToEx(hdc HDC, x, y int, lpPoint *POINT) bool {
	ret, _, _ := procMoveToEx.Call(
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(lpPoint)))

	return ret != 0
}


/////////////////////////////////////////////////
/////////////////////////////////////////////////


const (
	maxD = 15
	TIMER_ID = 1
	INTERVAL_V = 10
)



type _PTS struct{
	x int
	y int
	dx int
	dy int
}

var (
	maxX = 600
	maxY = 450
	pts [2]_PTS
	_pts [2]_PTS
	rect *RECT
	backColor COLORREF = 0x00ffffff
	// penF w32.HPEN
)

// func On_FormClose(arg *gform.EventArg) {
// 	w32.PostQuitMessage(0)
// }

func On_Paint(arg *gform.EventArg) {
	// hdc := w32.GetDC(arg.Sender().Handle())
	// hdcMemory := w32.CreateCompatibleDC(hdc)
	// hbmMemory := w32.CreateCompatibleBitmap(hdc, maxX, maxY)
	// w32.SelectObject(hdcMemory, (w32.HGDIOBJ)(unsafe.Pointer(hbmMemory)))

	// //w32.FillSolidRect(hdcMemory, 0, 0, maxX, maxY, w32.RGB(255, 255, 255))

	// pts[0], pts[1] = _pts[0], _pts[1]
	// for i := 0; i < 100; i++ {
	// 	Do_MovePoints()
	// 	if i == 0 {
	// 		_pts[0], _pts[1] = pts[0], pts[1]
	// 	}
	// 	w32.MoveToEx(hdcMemory, pts[0].x, pts[0].y, nil)
	// 	w32.LineTo(hdcMemory, pts[1].x, pts[1].y)
	// }
	// w32.BitBlt(hdc, 0, 0, maxX, maxY, hdcMemory, 0, 0, w32.SRCCOPY)
	// w32.DeleteObject((w32.HGDIOBJ)(unsafe.Pointer(hbmMemory)))
	// w32.DeleteDC(hdcMemory)

	// println(rect.Right)

	hwnd := arg.Sender().Handle()
	hdc := GetDC(hwnd)
	// w32.SetBkColor(hdc, backColor)
	// w32.FillSolidRect(hdc, 0, 0, maxX, maxY, backColor)
	// pen := w32.CreatePen(w32.PS_SOLID, 1, w32.RGB(128, 128, 128))
	// w32.SelectObject(hdc, (w32.HGDIOBJ)(unsafe.Pointer(penF)))
	pts[0], pts[1] = _pts[0], _pts[1]
	for i := 0; i < 100; i++ {
		Do_MovePoints()
		if i == 0 {
			_pts[0], _pts[1] = pts[0], pts[1]
		}
		MoveToEx(hdc, pts[0].x, pts[0].y, nil)
		LineTo(hdc, pts[1].x, pts[1].y)
	}
	ReleaseDC(hwnd, hdc)
}

func Do_InitPoints() {
	for i := 0; i < 2; i++ {
		_pts[i].x = rand.Intn(maxX)
		_pts[i].y = rand.Intn(maxY)
		_pts[i].dx = rand.Intn(maxD) - maxD / 2
		_pts[i].dy = rand.Intn(maxD) - maxD / 2
	}
}

func Do_MovePoints(){
	for i := 0; i < 2; i++ {
		nx := pts[i].x + pts[i].dx
		if nx >= 0 && nx < maxX {
			pts[i].x = nx
		} else {
			pts[i].dx = -pts[i].dx
		}
		ny := pts[i].y + pts[i].dy
		if ny >= 0 && ny < maxY {
			pts[i].y = ny
		} else {
			pts[i].dy = -pts[i].dy
		}
	}
}

func On_Timer(arg *gform.EventArg) {
	hwnd := arg.Sender().Handle()
	InvalidateRect(hwnd, nil, false)
	// println(_pts[0].x)
}

func main() {

	gform.Init()

	mf := gform.NewForm(nil)
	mf.SetSize(600, 450)
	mf.Center()

	// mf.Bind(w32.WM_CLOSE, On_FormClose)
	mf.Bind(WM_CLOSE, func(arg *gform.EventArg){
		PostQuitMessage(0)
	})

	mf.Bind(WM_SIZE, func(arg *gform.EventArg){
		hwnd := arg.Sender().Handle()
		rect = GetClientRect(hwnd)
		if rect.Right<=0 || rect.Bottom<=0 {
			return
		}
		maxX, maxY = int(rect.Right), int(rect.Bottom)
		dx, dy := maxX / 50, maxY / 50
		for i := 0; i < 2; i++ {
			_pts[i].dx = rand.Intn(dx)
			_pts[i].dy = rand.Intn(dy)
		}
		Do_InitPoints()
	})

	mf.OnPaint().Bind(On_Paint)
	mf.Bind(WM_TIMER, On_Timer)

	mf.Show()

	hwnd := mf.Handle()
	// penF = w32.CreatePen(w32.PS_SOLID, 1, backColor)

	SetTimer(hwnd, TIMER_ID, INTERVAL_V, 0)

	gform.RunMainLoop()
}
