package entities

import "log"

type SpawnState uint8

const (
	Counting SpawnState = iota
	Waiting
	Spawning
)

type Wave struct {
	Name  string
	Count int
	Enemy EnemyTypeId
	// Spawn rate in frames
	Delay int
}

// time in frames
type WaveSpawner struct {
	Waves             []Wave
	State             SpawnState
	nextWave          int
	timeBeetweenWaves int
	waveCountdown     int
	searchCountdown   int // TODO: 4:36 tutorial
}

func (w *WaveSpawner) Start() {
	w.waveCountdown = w.timeBeetweenWaves
}

func (w *WaveSpawner) Update() {
	if w.State == Waiting {

	}

	if w.waveCountdown <= 0 {
		if w.State != Spawning {
			w.SpawnWave(w.Waves[w.nextWave])
		}
	} else {
		w.waveCountdown -= 1
	}
}

func (w *WaveSpawner) EnemyIsAlive(enemies []GameObject) bool {
	return len(enemies) == 0
}

func (w *WaveSpawner) SpawnWave(wave Wave) {
	w.State = Spawning
	delay := wave.Delay

	for i := 0; i < wave.Count; i++ {
		if wave.Delay <= 0 {
			w.SpawnEnemy(wave.Enemy)
			wave.Delay = delay
		}

		wave.Delay -= 1
	}

	w.State = Waiting
}

func (w *WaveSpawner) SpawnEnemy(enemy EnemyTypeId) {
	log.Println("Enemy of type %d spawned", enemy)
}

// func (w *Wave) SpawnEnemies(scene Scene) {
// 	sceneObjects := *scene.GetObjects()
// 	portals := sceneObjects[PortalsObjectId]

// 	for enemyTypeId, quantity := range w.EnemyTypeQuantity {
// 		delay := w.SpawnDelay

// 		if enemyTypeId == BombheadTypeId {
// 			for i := 0; i < quantity; {
// 				delay -= 1
// 				if delay <= 0 {
// 					scene.AddObject(EnemiesObjectId, NewBombHead(
// 						GetRandomPortalPos(portals).Unpack(),
// 					))
// 					delay = w.SpawnDelay
// 				}

// 			}
// 		}

// 		if enemyTypeId == Shooter {
// 			for i := 0; i < quantity; {
// 				delay -= 1
// 				if delay <= 0 {
// 					scene.AddObject(EnemiesObjectId, NewShootyEnemy(
// 						GetRandomPortalPos(portals).Unpack(),
// 					))
// 					delay = w.SpawnDelay
// 				}
// 			}
// 		}
// 	}
// }

// func GetRandomPortalPos(portals []GameObject) *utils.Vec {
// 	randIndex := rand.Intn(len(portals))
// 	portal := portals[randIndex]

// 	return portal.GetCollider().GetPos()
// }
