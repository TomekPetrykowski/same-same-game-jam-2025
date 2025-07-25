package spritesheets

type SpritesheetId uint

const (
	PlayerSpritesheetId SpritesheetId = iota
	BombHeadSpritesheetId
)

var DB = map[SpritesheetId]*Spritesheet{
	PlayerSpritesheetId:   NewSpritesheet(2, 3, 15, 26),
	BombHeadSpritesheetId: NewSpritesheet(2, 3, 17, 27),
}
