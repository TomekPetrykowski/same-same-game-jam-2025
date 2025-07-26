package main

import (
	"fmt"
	"game/scenes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Scenes map[scenes.SceneId]scenes.Scene

type Game struct {
	scenes        map[scenes.SceneId]scenes.Scene
	activeSceneId scenes.SceneId
}

func NewGame() *Game {
	sceneMap := Scenes{
		// scenes.DebugSceneId:     scenes.NewDebugScene(),
		scenes.TestLevelSceneId:      scenes.NewTestLevelScene(),
		scenes.SortingSpritesSceneId: scenes.NewSortingSpritesScene(),
		scenes.StartSceneId:          scenes.NewStartScene(),
	}

	activeSceneId := scenes.StartSceneId

	sceneMap[activeSceneId].FirstLoad()

	return &Game{
		sceneMap,
		activeSceneId,
	}

}

func (g *Game) Update() error {
	nextSceneId := g.scenes[g.activeSceneId].Update()

	// if key 1 or key 2 are pressed scene is changing
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		nextSceneId = g.scenes[5].Update()
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
		nextSceneId = g.scenes[6].Update()
	}

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
	return 640, 360
}

func (g *Game) Set() {
	fmt.Println("Game is running!")
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Pothead")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}
