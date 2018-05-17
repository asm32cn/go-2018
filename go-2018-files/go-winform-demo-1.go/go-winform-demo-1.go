// go-winform-demo-1.go
package main

import (
	// "github.com/lxn/go-winapi"
	"go-winapi-master"
	"syscall"
	"strconv"
	"unsafe"
)

var (
	orainWndProc
	winapi.HWND
)

const(
	winWidth int32 = 500
	winHeight int32 = 300
)

func _TEXT(_str string) *uint16{
	return syscall.StringToUTF16Ptr(_str)
}

func _toString(_n int32) string{
	return strconv.Itoa(int(_n))
}

func WndProc(hwnd winapi.HWND,msg uint32,wparam uintptr,lparam uintptr) uintptr{
	return winapi.CallWindowProc(uintptr(orainWndProc),hwnd,msg,wparam,lparam)
}

func main(){
	var message winapi.MSG
	var hwnd winapi.HWND
	var wproc uintptr hwnd = winapi.CreateWindowEx(
		winapi.WS_EX_CLIENTEDGE,
		_TEXT("EDIT"),
		_TEXT("Hello World"),
		winapi.WS_OVERLAPPEDWINDOW,
		(winapi.GetSystemMetrics(winapi.SM_CXSCREEN)-winWidth)>>1,
		(winapi.GetSystemMetrics(winapi.SM_CYSCREEN)-winHeight)>>1,
		winWidth,
		winHeight,
		0,
		0,
		winapi.GetModuleHandle(nil),
		unsafe.Pointer(nil)
	)
	
	wproc = syscall.NewCallback(WndProc)
	orainWndProc = winapi.HWND(winapi.SetWindowLong(hwnd,winapi.GWL_WNDPROC, int32(wproc)))
	winapi.ShowWindow(hwnd,winapi.SW_SHOW)
	
	for{
		if winapi.GetMessage(&message,0,0,0) == 0{break}
		winapi.TranslateMessage(&message)
		winapi.DispatchMessage(&message)
	}
}