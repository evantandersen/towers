package main

import (
    "os"
    "io/ioutil"
    "encoding/json"
	"image"
    _ "image/png"
    gl "github.com/go-gl/gl"
)

type ImageInfo struct {
    Frame Rect    
}

type Size struct {
    W int 
    H int
}

type Meta struct {
    Image string 
    Size Size 
}

type ImageData struct {
    Frames map[string]ImageInfo
    Meta Meta
}

var Coords map[string]Rect

func loadTextures(infoPath string) error {
    //load the json
	jsonData, err := ioutil.ReadFile(infoPath)
	if err != nil {
		return err
	}
    
    var info ImageData
    err = json.Unmarshal(jsonData, &info)
    if err != nil {
        return err
    }
    
    Coords = make(map[string]Rect)
    w := float64(info.Meta.Size.W)
    h := float64(info.Meta.Size.H)
    for k, v := range info.Frames {
        Coords[k] = Rect{
            X:v.Frame.X/w,
            Y:v.Frame.Y/h,
            W:v.Frame.W/w,
            H:v.Frame.H/h,
        }
    }
    
    r, err := os.Open(info.Meta.Image)
	if err != nil {
		return err
	}
	defer r.Close()
	
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	rgba := image.NewRGBA(img.Bounds())

    width := img.Bounds().Max.X
    height := img.Bounds().Max.Y

	for j := 0; j < height; j++ {
		for k := 0; k < width; k++ {
			rgba.Set(k, j, img.At(k,j))
		}
	}
    
	texture := gl.GenTexture()
	texture.Bind(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
    gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, rgba.Pix)
           
	return nil
}

type Point struct {
    X, Y float64
}

type Path []Point

type Level struct {
    BackgroundImage string
    Waves []*Spawner
    Path Path
}

func loadLevel(name string) (result *Level, err error) {
    //load the json
	jsonData, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}
    
    var level Level
    err = json.Unmarshal(jsonData, &level)
    if err != nil {
        return
    }
    //TODO sanity checks
    
    result = &level
    return
}