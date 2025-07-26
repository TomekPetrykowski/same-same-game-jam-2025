package ui

import (
	"bytes"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Font struct {
	fontFaceSource *text.GoTextFaceSource
}

type TextToRender struct {
	Text                    string
	X, Y, Size, LineSpacing float64
	Color                   color.Color
}

func NewFont() *Font {
	fontBytes, err := os.ReadFile("assets/fonts/tiny5.ttf")
	if err != nil {
		log.Fatal(err)
	}

	source, err := text.NewGoTextFaceSource(bytes.NewReader(fontBytes))
	if err != nil {
		log.Fatal(err)
	}

	return &Font{fontFaceSource: source}
}

func (f *Font) Render(screen *ebiten.Image, data *TextToRender) {
	opts := text.DrawOptions{}
	opts.GeoM.Translate(data.X, data.Y)
	opts.ColorScale.ScaleWithColor(data.Color)
	opts.LineSpacing = data.LineSpacing
	text.Draw(screen, data.Text, &text.GoTextFace{
		Source: f.fontFaceSource,
		Size:   data.Size,
	}, &opts)
}

func (f *Font) MeasureText(data *TextToRender) (float64, float64) {
	return text.Measure(data.Text, &text.GoTextFace{
		Source: f.fontFaceSource,
		Size:   data.Size,
	}, data.Size*data.LineSpacing)
}
