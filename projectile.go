package main

type Projectile struct {
    Graphic
    Damage float64
    Speed float64
    target *Mob
}

func (p *Projectile)Update(state *GameState) {
    if p.MoveTowardsPoint(p.target.X, p.target.Y, p.Speed) > 0 {
        p.target.Health -= p.Damage
        state.removeID(p.ID)
    }
    p.AimTowardsPoint(p.target.X, p.target.Y)
}
