package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/tropolab/wx/d3"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	winX      = 200
	winY      = 200
	winWidth  = 720
	winHeight = 200
	nullStr   = "\x00"
)

var vertexShaderSource = `
#version 330 core

layout (location=0) in vec3 aPos;

void main()
{
	gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
}` + nullStr

var fragmentShaderSource = `
#version 330 core

out vec4 FragColor;

void main()
{
	FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
}` + nullStr

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat(nullStr, int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func newProgram(vertexShaderSrc, fragmentShaderSrc string) (uint32, error) {

	// compile vertextshader
	vertexShader, err := compileShader(vertexShaderSrc, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	// compile fragmentshader
	fragmentShader, err := compileShader(fragmentShaderSrc, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	// create a program object
	program := gl.CreateProgram()

	// attach vertexshader to program object
	gl.AttachShader(program, vertexShader)

	// attach fragmentshader to program object
	gl.AttachShader(program, fragmentShader)

	// link the program object
	gl.LinkProgram(program)

	// verify program
	var status int32

	if gl.GetProgramiv(program, gl.LINK_STATUS, &status); status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat(nullStr, int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	// delete vertexshader
	gl.DeleteShader(vertexShader)

	// delete fragmentshader
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func main() {

	polygon := d3.Polygon{
		A: d3.Vertex{X: -0.5, Y: -0.5, Z: 0.0},
		B: d3.Vertex{X: 0.5, Y: -0.5, Z: 0.0},
		C: d3.Vertex{X: 0.0, Y: 0.5, Z: 0.0},
	}
	polygons := []d3.Polygon{}
	polygons = append(polygons, polygon)

	mesh := d3.Mesh{Polygons: polygons}

	if err := mesh.ToObj(); err != nil {
		log.Printf("error rendering 3d mesh")
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalln("failed to initialize SDL2:", err)
		panic(err)
	}
	defer sdl.Quit()

	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	window, err := sdl.CreateWindow("TropoLab Weather", winX, winY, winWidth, winHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	window.GLCreateContext()

	// Initialize gl
	if err := gl.Init(); err != nil {
		log.Fatalln("failed to initialize OpenGL:", err)
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// Configure the vertex and fragment shaders
	program, err := newProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		panic(err)
	}

	// render first polygon from mesh
	vertices := mesh.Polygons[0].Render()

	// bind vertice array
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	// bind vertice buffer
	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)

	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// tell gl how vertices are laid out
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		// set black background
		gl.ClearColor(0.0, 0.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// tell gl use the shader program
		gl.UseProgram(program)

		// bind the vertex array
		gl.BindVertexArray(VAO)

		// render the vertex array as triangles
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		// swap the window buffers
		window.GLSwap()
	}

}
