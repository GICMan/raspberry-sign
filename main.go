package main

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// void fullscreen(void *win)
// {
// 	gtk_window_fullscreen(GTK_WINDOW(win));
// }
import "C"
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

	// Use custom C function to fullscreen
	C.fullscreen(w.Window())

	w.Navigate("https://google.com")
	w.Run()
}
