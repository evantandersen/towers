package main


type Updateable interface {
    Update(*GameState)
}

type GameState struct {
    Renderables map[id]Renderable
    Entities map[id]Updateable
    Targets map[id]*Mob
    Lives int
    Time float64
    Level *Level
}

func NewGameState(level *Level) *GameState {
    var state GameState
    state.Lives = 20
    state.Level = level
    state.Renderables = make(map[id]Renderable)
    state.Entities = make(map[id]Updateable)
    state.Targets = make(map[id]*Mob)
    
    //load the bg
    a := &Graphic{}
    a.ID = NewID()
    a.W = 1.0
    a.H = 1.0
    a.X = 0.5
    a.Y = 0.5
    a.Z = 0.0
    a.TextureName = state.Level.BackgroundImage
    state.Renderables[a.ID] = a
    
    //add the mob spawners
    for _, wave := range state.Level.Waves {
        id := NewID()
        wave.ID = id
        state.Entities[id] = wave
    }
    
    return &state
}

func (s *GameState)removeID(id id) {
    delete(s.Renderables, id)
    delete(s.Entities, id)
    delete(s.Targets, id)
}

