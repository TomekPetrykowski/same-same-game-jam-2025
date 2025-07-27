package entities

import (
	utils "game/utils/math"
	"math/rand"
)

type Wave struct {
	EnemyTypeQuantity map[EnemyTypeId]int
	// Spawn delay in frames
	SpawnDelay int
}

func (w *Wave) SpawnEnemies(scene Scene) {
	sceneObjects := *scene.GetObjects()
	portals := sceneObjects[PortalsObjectId]

	for enemyTypeId, quantity := range w.EnemyTypeQuantity {
		delay := w.SpawnDelay

		if enemyTypeId == BombheadTypeId {
			for i := 0; i < quantity; {
				delay -= 1
				if delay <= 0 {
					scene.AddObject(EnemiesObjectId, NewBombHead(
						GetRandomPortalPos(portals).Unpack(),
					))
					delay = w.SpawnDelay
				}

			}
		}

		if enemyTypeId == Shooter {
			for i := 0; i < quantity; {
				delay -= 1
				if delay <= 0 {
					scene.AddObject(EnemiesObjectId, NewShootyEnemy(
						GetRandomPortalPos(portals).Unpack(),
					))
					delay = w.SpawnDelay
				}
			}
		}
	}
}

func GetRandomPortalPos(portals []GameObject) *utils.Vec {
	randIndex := rand.Intn(len(portals))
	portal := portals[randIndex]

	return portal.GetCollider().GetPos()
}
