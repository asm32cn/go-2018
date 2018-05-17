// go-dialog-demo-2.go

package main

import (
    "github.com/AllenDang/gform"
    "github.com/AllenDang/w32"
)

func main() {
    gform.Init()

    mf := gform.NewForm(nil)
    mf.SetSize(600, 450)
    mf.Center()
    mf.Show()

    mf.Bind(w32.WM_CLOSE, func(arg *gform.EventArg){
    	w32.PostQuitMessage(0)
    })

    gform.RunMainLoop()
}