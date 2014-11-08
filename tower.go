package main

import (
    "math"
)

type Tower struct {
    Graphic
    Interval float64
    Range float64
    Ammo Projectile
    lastShot float64
}

func (t *Tower)Update(state *GameState) {
    if state.Time > (t.Interval + t.lastShot) && len(state.Targets) > 0 {
        minDistance := 9999999.0
        var closest *Mob
        for _,target := range state.Targets {
            if target.futureHealth <= 0 {
                continue
            }
            diffX := t.X - target.X
            diffY := t.Y - target.Y
            distance := math.Sqrt(diffX*diffX + diffY*diffY)
            if distance < minDistance {
                minDistance = distance
                closest = target
            }
        }
        if minDistance > t.Range {
            return
        }
        
        var bullet Projectile
        bullet.ID = NewID()
        bullet.Speed = 0.01
        bullet.X = t.X
        bullet.Y = t.Y
        bullet.W = 0.08
        bullet.H = 0.10
        bullet.Damage = 2
        bullet.TextureName = "fireball.png"
        bullet.Z = 3.0
        bullet.target = closest
        closest.futureHealth -= 2
        state.Renderables[bullet.ID] = &bullet
        state.Entities[bullet.ID] = &bullet
    
        t.lastShot = state.Time
    }    
}
