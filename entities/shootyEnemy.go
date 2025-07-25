package entities

type ShootyEnemy struct {
	Entity
	Cooldown float64
	// Speed     float64
	// Direction v.Vec
}

func (e *ShootyEnemy) Update(scene Scene) {
	playerPos := (*(*scene.GetObjects())["player"][0].GetCollider()).GetPos()
	// e.Pos.Add(e.Pos.DirectionTo(*playerPos).Multiplied(0.5))
	direction := e.Collider.GetPos().DirectionTo(*playerPos).Normalized()

	e.Cooldown -= 0.01
	if e.Cooldown <= 0 {
		// print("shoot")
		e.Cooldown = 1
		// scene.GetObjects()["enemyProjectiles"][] = append((*scene.GetObjects())["enemyProjectiles"], NewBasicProjectile((*e.GetCollider()).GetPos().X, (*e.GetCollider()).GetPos().Y, 1, direction))
		scene.AddObject("enemyProjectiles", NewBasicProjectile((*e.GetCollider()).GetPos().X, (*e.GetCollider()).GetPos().Y, 1, direction))
	}

}

func NewShootyEnemy(x, y float64) *ShootyEnemy {
	ent := NewEntity(NewCircle(x, y, 10), nil)
	return &ShootyEnemy{Entity: *ent, Cooldown: 1}
}
