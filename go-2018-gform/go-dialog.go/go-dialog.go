package main

import (
	//"fmt"
	"github.com/AllenDang/gform"
	"github.com/AllenDang/w32"
)

var (
	edt *gform.Edit
)

func onclick(arg *gform.EventArg) {
	gform.MsgBox(arg.Sender(), "消息", "Hello world", w32.MB_OK | w32.MB_ICONINFORMATION)
	// println("aaa")
}

func main() {
	gform.Init()

	mainWindow := gform.NewForm(nil)
	// mainWindow.SetPos(300, 100)
	mainWindow.SetSize(500, 300)
	mainWindow.Center()
	mainWindow.SetCaption("Controls Demo")
	mainWindow.Bind(w32.WM_CLOSE, func(arg *gform.EventArg){
		w32.PostQuitMessage(0);
	})

	btn := gform.NewPushButton(mainWindow)
	btn.SetPos(10, 10)
	btn.OnLBUp().Bind(onclick)

	mainWindow.Show()

	gform.RunMainLoop()
}
