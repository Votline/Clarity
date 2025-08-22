package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

func Setup() uint32 {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	pg := gl.CreateProgram()
	
	shaders := attachShaders(pg)

	gl.LinkProgram(pg)
	gl.UseProgram(pg)

	detachShaders(pg, shaders)

	gl.ClearColor(0.0, 0.0, 0.0, 0.7)

	return pg
}
