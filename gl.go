package main

import (
    "fmt"
    "log"
    "sort"
    gl "github.com/go-gl/gl"
)

func checkGLError() {
	glErr := gl.GetError()
	if glErr != 0 {
		log.Printf("GL Error Code: %d\n", int(glErr))
		panic("stack trace")
	}
}

type Polygon struct {
    Vertices []float32
    Z float32
    ID id
}

type ByZ []Polygon

func (a ByZ) Len() int           { return len(a) }
func (a ByZ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a ByZ) Less(i, j int) bool {
    if a[i].Z == a[j].Z {
        return a[j].ID[0] < a[i].ID[0]
    }
    return a[i].Z < a[j].Z 
}

func drawRenderables(toRender map[id]Renderable) {
    polygons := make([]Polygon, len(toRender))
    i := 0
    for id,renderable := range toRender {
        polygons[i].Vertices, polygons[i].Z = renderable.Render()
        polygons[i].ID = id
        i++
    }
    sort.Sort(ByZ(polygons))
    
    var vertices []float32
    for _, polygon := range polygons {
        vertices = append(vertices, polygon.Vertices...)
    }
    drawVertices(vertices)
}

func drawVertices(vertices []float32) {
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), vertices, gl.DYNAMIC_DRAW)
    gl.DrawArrays(gl.TRIANGLES, 0, len(vertices)/4)
}


func start() error {
	err := ShaderInit() 
	if err != nil {
		return err
	}
    
    gl.Enable(gl.BLEND)
    gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
    
	vao := gl.GenVertexArray()
	vao.Bind()
    
	vertexBuf := gl.GenBuffer()
	vertexBuf.Bind(gl.ARRAY_BUFFER)
            
    err = loadTextures("spirte_data.json")
    if err != nil {
        return err
    }
            
	program, err := loadProgram("vertex.glsl", "fragment.glsl")
	if err != nil {
        return fmt.Errorf("Failed to load preprocessing program: %v", err)
	}
	
    program.Use()
    
	posLocation := program.GetAttribLocation("position")
	checkGLError()	
	posLocation.AttribPointer(2, gl.FLOAT, false, 4*4, nil)
	posLocation.EnableArray()
        
	oldlocation := program.GetUniformLocation("img")
	oldlocation.Uniform1i(0)
    
	texLocation := program.GetAttribLocation("TexCoord")
	checkGLError()	
	texLocation.AttribPointer(2, gl.FLOAT, false, 4*4, uintptr(2*4))
	texLocation.EnableArray()    
    checkGLError()	
    
    return nil
}