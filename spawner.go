package main

import (
)

type Spawner struct {
    Entity
    Mob Mob
    Count int
    Interval float64
    Start float64
    index int
}

func (s *Spawner)Update(state *GameState) {
    nextMobTime := float64(s.index)*s.Interval + s.Start
    if state.Time > nextMobTime {
        newMob := s.Mob
        newMob.X = state.Level.Path[0].X
        newMob.Y = state.Level.Path[0].Y
        newMob.futureHealth = newMob.Health
        newMob.ID = NewID()
        newMob.pathIndex = 1
        state.Entities[newMob.ID] = &newMob
        state.Renderables[newMob.ID] = &newMob
        state.Targets[newMob.ID] = &newMob
        s.index++
        if s.index >= s.Count {
            state.removeID(s.ID)
        }
    }
}

