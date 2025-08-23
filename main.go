package main

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"

	"Clarity/internal/ui"
	"Clarity/internal/core"
	"Clarity/internal/render"
	"Clarity/internal/config"
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
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("Config getting error: \n%v", err)
	}

	for {
		win := ui.PrimaryWindow(cfg.AlignX, cfg.AlignY, cfg.WinW, cfg.WinH)
		pg, ofl := render.Setup(cfg.BackC, cfg.TextC)

		hv := ui.CreateView(pg, ofl)

		done := false
		var remaining string
		cmdChan := make(chan string)
		timerChan := make(chan string)
		go core.Timer(win, timerChan, cfg.Sound, &done, cmdChan)

		gl.LineWidth(2.0)
		glfw.SwapInterval(1)
		for !win.ShouldClose() {
			select {
			case cmd := <- cmdChan:
				switch cmd {
				case "close":
					win.SetShouldClose(true)
				case "show":
					win.Show()
				}
			default:
			}

			gl.Clear(gl.COLOR_BUFFER_BIT)

			select {
			case remaining = <- timerChan:
			default:
			}
			hv.Render(remaining)

			win.SwapBuffers()
			glfw.PollEvents()
		}
		win.Destroy()
	}
}
