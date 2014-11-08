package main

import (
    "math"
)

type Rect struct {
    X float64
    Y float64
    W float64
    H float64
}

type Graphic struct {
    Entity
    Rect
    TextureName string
    Z float64
    Angle float64
}

type Renderable interface {
    Render() ([]float32, float32)
}

func (g *Graphic)Render() (vertices []float32, z float32) {    
    xStart := -g.W/2
    yStart := -g.H/2
    xEnd := g.W/2 
    yEnd := g.H/2

    var x0,x1,x2,x3,y0,y1,y2,y3 float32
    if g.Angle != 0 {
        sinTheta, cosTheta := math.Sincos(g.Angle)
        
        x0 = float32(xStart*cosTheta - yStart*sinTheta)
        y0 = float32(xStart*sinTheta + yStart*cosTheta)
        
        x1 = float32(xStart*cosTheta - yEnd*sinTheta)
        y1 = float32(xStart*sinTheta + yEnd*cosTheta)

        x2 = float32(xEnd*cosTheta - yEnd*sinTheta)
        y2 = float32(xEnd*sinTheta + yEnd*cosTheta)

        x3 = float32(xEnd*cosTheta - yStart*sinTheta)
        y3 = float32(xEnd*sinTheta + yStart*cosTheta)
    } else {
        x0 = float32(xStart)
        x1 = float32(xStart)
        x2 = float32(xEnd)
        x3 = float32(xEnd)
        
        y0 = float32(yStart)
        y1 = float32(yEnd)
        y2 = float32(yEnd)
        y3 = float32(yStart)
    }
    
    x0 += float32(g.X)
    x1 += float32(g.X)
    x2 += float32(g.X)
    x3 += float32(g.X)

    y0 += float32(g.Y)
    y1 += float32(g.Y)
    y2 += float32(g.Y)
    y3 += float32(g.Y)
    

    texRect := Coords[g.TextureName]
    texXStart := float32(texRect.X)
    texYStart := float32(texRect.Y)
    texXEnd := float32(texRect.X + texRect.W)
    texYEnd := float32(texRect.Y + texRect.H)
        
    vertices = []float32{
        // bottom left triangle
        x0, y0, texXStart, texYStart, x1, y1, texXStart, texYEnd, x3, y3, texXEnd, texYStart,

        // upper right triangle   
        x1, y1, texXStart, texYEnd, x2, y2, texXEnd, texYEnd, x3, y3, texXEnd, texYStart,
    }
    z = float32(g.Z)
    return
}

func (g *Graphic)MoveTowardsPoint(X, Y, delta float64) float64 {
    diffX := X - g.X
    diffY := Y - g.Y
    hyp := math.Sqrt(float64(diffX*diffX + diffY*diffY))
    distance := delta
    if distance >= hyp {
        distance = hyp
    }
    g.X += (diffX/hyp) * distance
    g.Y += (diffY/hyp) * distance
    
    return delta - distance
}

func (g *Graphic)AimTowardsPoint(X, Y float64) {
    diffX := X - g.X
    diffY := Y - g.Y
    
    // +pi/2 to make 0,1 = 0 degrees
    g.Angle = math.Atan2(diffY, diffX) + math.Pi/2
}