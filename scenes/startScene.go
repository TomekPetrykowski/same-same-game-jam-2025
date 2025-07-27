package scenes

import (
	"game/entities"
	"game/ui"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type StartScene struct {
	loaded  bool
	objects entities.ObjectsMap
	font    *ui.Font
}

func NewStartScene() *StartScene {
	return &StartScene{
		loaded:  false,
		objects: nil,
		font:    nil,
	}
}

func (s *StartScene) AddObject(entities.SceneObjectId, entities.GameObject) {

}

func (s *StartScene) GetObjects() *entities.ObjectsMap {
	return &s.objects
}

func (s *StartScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
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

	startGame := ui.TextToRender{
		Text:  "Press Enter to start game",
		Color: color.RGBA{200, 197, 150, 255},
		Size:  16,
		Y:     gameName.Y + 50,
	}

	descWidth, _ := s.font.MeasureText(&startGame)
	startGame.X = (640 / 2) - (descWidth / 2)

	s.font.Render(screen, &gameNameShadow)
	s.font.Render(screen, &gameName)
	s.font.Render(screen, &startGame)
}

func (s *StartScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return LevelSceneId
	}

	return StartSceneId
}

func (s *StartScene) FirstLoad() {
	s.font = ui.NewFont()
	s.loaded = true
}

func (s *StartScene) IsLoaded() bool {
	return s.loaded
}

func (s *StartScene) OnEnter() {

}

func (s *StartScene) OnExit() {

}

var _ Scene = (*StartScene)(nil)
var _ entities.Scene = (*StartScene)(nil)
