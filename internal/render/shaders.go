package render

import (
	"log"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const vertexShaderSource = `
#version 410 core

layout (location = 0) in vec3 aPos;
uniform float xOffset;

void main() {
	gl_Position = vec4(aPos.x + xOffset, aPos.y, aPos.z, 1.0);
}` + "\x00"

const fragmentShaderSource = `
#version 410 core

out vec4 FragColor;

void main() {
	FragColor = vec4(1.0, 1.0, 1.0, 1.0);
}` + "\x00"

func attachShaders(pg uint32) []uint32 {
	vertexShader := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	fragmentShader := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)

	gl.AttachShader(pg, vertexShader)
	gl.AttachShader(pg, fragmentShader)

	return []uint32{vertexShader, fragmentShader}
}
func compileShader(source string, shaderType uint32) uint32 {
	shader := gl.CreateShader(shaderType)
	
	csource, free := gl.Strs(source)
	defer free()

	gl.ShaderSource(shader, 1, csource, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logs := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logs))
		
		gl.DeleteShader(shader)
		log.Fatalf("Compile shader error: \n%v", logs)
	}

	return shader
}

func detachShaders(pg uint32, shaders []uint32) {
	for _, shader := range shaders {
		gl.DetachShader(pg, shader)
		gl.DeleteShader(shader)
	}
}
