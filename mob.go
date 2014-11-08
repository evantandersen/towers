package main

import (
)

type Mob struct {
    Graphic
    Speed float64
    Bounty float64
    Health float64
    futureHealth float64
    pathIndex int
}

func (m *Mob)Update(state *GameState) {
    if m.Health <= 0 {
        state.removeID(m.ID)
        return
    }
    
    deltaLeft := m.Speed
    for deltaLeft > 0 {
        nextPoint := state.Level.Path[m.pathIndex]
        deltaLeft = m.MoveTowardsPoint(nextPoint.X, nextPoint.Y, deltaLeft)
        if deltaLeft > 0 {
            m.pathIndex++
            if m.pathIndex >= len(state.Level.Path) {
                state.Lives--
                state.removeID(m.ID)
                break
            }
        }
    }
}
