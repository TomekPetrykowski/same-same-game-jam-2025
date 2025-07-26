package spritesheets

type SpritesheetId uint

const (
	PlayerSpritesheetd SpritesheetId = iota
	BombheadSpritesheetd
	BirdmanSpritesheetd
	BoneyknightSpritesheetId
)

var DB = map[SpritesheetId]*Spritesheet{
	PlayerSpritesheetId:      NewSpritesheet(2, 3, 15, 26),
	BombheadSpritesheetId:    NewSpritesheet(2, 3, 17, 27),
	BirdmanSpritesheetId:     NewSpritesheet(2, 2, 51, 57),
	BoneyknightSpritesheetId: NewSpritesheet(6, 3, 54, 52),
}
