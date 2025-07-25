package entities

type ShootyEnemy struct {
	*Entity
	Cooldown float64
	// Speed     float64
	// Direction v.Vec
}

func NewShootyEnemy(x, y float64) *ShootyEnemy {
	ent := NewEntity(NewCircle(x, y, 10), nil)
	return &ShootyEnemy{Entity: ent, Cooldown: 1}
}

func (se *ShootyEnemy) Update(scene Scene) {
	se.ShootPlayer(scene)
}

func (se *ShootyEnemy) ShootPlayer(scene Scene) {
	sceneObjects := *scene.GetObjects()
	player := sceneObjects[PlayerObjectId][0]
	playerPos := player.GetCollider().GetPos()
	enemyPos := se.GetCollider().GetPos()
	direction := se.Collider.GetPos().DirectionTo(*playerPos).Normalized()

	se.Cooldown -= 0.01
	if se.Cooldown <= 0 {
		se.Cooldown = 1
		scene.AddObject(
			EnemyProjectilesObjectId,
			NewBasicProjectile(enemyPos.X, enemyPos.Y, 1, direction),
		)
	}
}
