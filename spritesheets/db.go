package spritesheets

type SpritesheetId uint

const (
	PlayerSpritesheedId SpritesheetId = iota
	BombheadSpritesheedId
	BirdmanSpritesheedId
	BoneyknightSpritesheedId
)

var DB = map[SpritesheetId]*Spritesheet{
	PlayerSpritesheedId:      NewSpritesheet(2, 3, 15, 26),
	BombheadSpritesheedId:    NewSpritesheet(2, 3, 17, 27),
	BirdmanSpritesheedId:     NewSpritesheet(2, 2, 51, 57),
	BoneyknightSpritesheedId: NewSpritesheet(6, 3, 54, 52),
}
