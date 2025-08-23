package main

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"

	"Clarity/internal/ui"
	"Clarity/internal/core"
	"Clarity/internal/render"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalf("GLFW init error: \n%v", err)
	}
	defer glfw.Terminate()
	
	if err := gl.Init(); err != nil {
		log.Fatalf("OpenGL init error: \n%v", err)
	}
	log.SetFlags(log.Lshortfile)

	for {
		win := ui.PrimaryWindow()
		pg, ofl := render.Setup()

		hv := ui.CreateView(pg, ofl)

		done := false
		var remaining string
		timerChan := make(chan string)
		go core.Timer(win, timerChan, &done)

		gl.LineWidth(2.0)
		glfw.SwapInterval(1)
		for !win.ShouldClose() {
			gl.Clear(gl.COLOR_BUFFER_BIT)

			remaining = <-timerChan
			hv.Render(remaining)

			win.SwapBuffers()
			glfw.PollEvents()
		}
		win.Destroy()
	}
}
