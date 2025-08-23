package ui

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const windowWidth = 150
const windowHeight = 70

type element interface{
	create() element
	getVtc() []float32
	getVtq() int32
	setNum(int32)
}
type digit struct {
	vtc []float32
	vtq int32
	num rune
}

func PrimaryWindow() *glfw.Window {
	glfw.WindowHint(glfw.RefreshRate, 60)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.DoubleBuffer, glfw.True)
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
	
	win, err := glfw.CreateWindow(windowWidth, windowHeight, "Clarity", nil, nil)
	if err != nil {
		log.Fatalf("Create window error: \n%v", err)
	}

	win.MakeContextCurrent()
	win.SetAttrib(glfw.Floating, 1)

	vidMode := glfw.GetPrimaryMonitor().GetVideoMode()
	win.SetPos(vidMode.Width-160, vidMode.Height-1075)
	win.Hide()

	return win
}
