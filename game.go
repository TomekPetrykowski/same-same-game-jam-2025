package main

import (
	"fmt"
	"game/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scenes map[scenes.SceneId]scenes.Scene

type Game struct {
	scenes        map[scenes.SceneId]scenes.Scene
	activeSceneId scenes.SceneId
}

func NewGame() *Game {
	sceneMap := Scenes{
		scenes.DebugSceneId:     scenes.NewDebugScene(),
		scenes.TestLevelSceneId: scenes.NewTestLevelScene(),
	}

	activeSceneId := scenes.TestLevelSceneId

	sceneMap[activeSceneId].FirstLoad()

	return &Game{
		sceneMap,
		activeSceneId,
	}

}

func (g *Game) Update() error {
	nextSceneId := g.scenes[g.activeSceneId].Update()

	if nextSceneId == scenes.ExitSceneId {
		g.scenes[g.activeSceneId].OnExit()
		return ebiten.Termination
	}

	if nextSceneId != g.activeSceneId {
		nextScene := g.scenes[nextSceneId]
		// if not loaded then load in
		if !nextScene.IsLoaded() {
			nextScene.FirstLoad()
		}

		nextScene.OnEnter()
		g.scenes[g.activeSceneId].OnExit()
	}

	g.activeSceneId = nextSceneId

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scenes[g.activeSceneId].Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) Set() {
	fmt.Println("Game is running!")
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
}
