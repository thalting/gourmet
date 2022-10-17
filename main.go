package main

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
import "C"

import (
	"flag"
	"os"
	"os/exec"
)

func main() {
	var rev C.int
	var nc C.uint
	var win C.Window
	var root C.Window
	var ch *C.Window

	if len(os.Args) == 1 {
		return
	}

	var shellalias bool
	flag.BoolVar(&shellalias, "s", false, "Shell aliases")
	flag.Parse()

	dpy := C.XOpenDisplay(nil)
	defer C.XCloseDisplay(dpy)
	if dpy == nil {
		panic("Can't open display")
	}

	C.XGetInputFocus(dpy, &win, &rev)

	if rev == C.RevertToParent {
		C.XQueryTree(dpy, win, &root, &win, &ch, &nc)
	}

	C.XUnmapWindow(dpy, win)
	C.XFlush(dpy)

	if shellalias == false {
		exec.Command(os.Args[1], os.Args[2:]...).Run()
	} else {
		cmd := []string{
			os.Getenv("SHELL"),
			"-i",
			"-c",
		}
		cmd = append(cmd, os.Args[2:]...)
		exec.Command(cmd[0], cmd[1:]...).Run()
	}

	C.XMapWindow(dpy, win)
}
