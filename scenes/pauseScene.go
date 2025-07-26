package scenes

import (
	"game/entities"
	"game/ui"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PauseScene struct {
	loaded  bool
	objects entities.ObjectsMap
	font    *ui.Font
}

func NewPauseScene() *PauseScene {
	return &PauseScene{
		loaded:  false,
		objects: nil,
		font:    nil,
	}
}

func (s *PauseScene) AddObject(entities.SceneObjectId, entities.GameObject) {

}

func (s *PauseScene) GetObjects() *entities.ObjectsMap {
	return &s.objects
}

func (s *PauseScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 30, 10})
	gameName := ui.TextToRender{
		Text:  "Pothead",
		Color: color.RGBA{249, 246, 199, 255}, // assets color
		Size:  32,
		Y:     100,
	}

	nameWidth, _ := s.font.MeasureText(&gameName)
	// game dims are 640x360
	gameName.X = (640 / 2) - (nameWidth / 2)

	gameNameShadow := ui.TextToRender{
		Text:  "Pothead",
		Color: color.RGBA{100, 97, 50, 255},
		Size:  32,
		X:     gameName.X + 2,
		Y:     gameName.Y + 2,
	}

	desc := ui.TextToRender{
		Text:        "Game is paused\nPress P to unpause\nPress R to restart game and press Q to exit",
		Color:       color.RGBA{200, 197, 150, 255},
		Size:        16,
		Y:           gameName.Y + 50,
		LineSpacing: 20,
	}

	descWidth, _ := s.font.MeasureText(&desc)
	desc.X = (640 / 2) - (descWidth / 2)

	s.font.Render(screen, &gameNameShadow)
	s.font.Render(screen, &gameName)
	s.font.Render(screen, &desc)
}

func (s *PauseScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		return LevelSceneId
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return ExitSceneId
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		return StartSceneId
	}

	return PauseSceneId
}

func (s *PauseScene) FirstLoad() {
	s.font = ui.NewFont()
	s.loaded = true
}

func (s *PauseScene) IsLoaded() bool {
	return s.loaded
}

func (s *PauseScene) OnEnter() {

}

func (s *PauseScene) OnExit() {

}

var _ Scene = (*PauseScene)(nil)
var _ entities.Scene = (*PauseScene)(nil)
