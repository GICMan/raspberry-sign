package main

import (
	"github.com/webview/webview"
)

var w webview.WebView

func main() {
	// Create new webview window
	w = webview.New(false)
	defer w.Destroy()
	w.SetTitle("piTimer")
	w.SetSize(480, 320, webview.HintNone)

	w.Navigate("https://google.com")
	w.Run()
}
