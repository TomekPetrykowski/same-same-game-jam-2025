package spritesheets

import "image"

type Spritesheet struct {
	WidthInTiles  int
	HeightInTiles int
	TileWidth     int
	TileHeight    int
}

func (s *Spritesheet) Rect(index int) image.Rectangle {
	x := (index % s.WidthInTiles) * s.TileWidth
	y := (index / s.WidthInTiles) * s.TileHeight

	return image.Rect(
		x, y, x+s.TileWidth, y+s.TileHeight,
	)
}

// Returns pointer to a new spritesheet data
//
// 	wit: width in tiles
// 	hit: height in tiles
//	tw: tile width
//	th: tile height
func NewSpritesheet(wit, hit, tw, th int) *Spritesheet {
	return &Spritesheet{
		wit, hit, tw, th,
	}
}
