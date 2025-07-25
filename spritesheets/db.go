package spritesheets

type SpritesheetId uint

const (
	PlayerSpritesheedId SpritesheetId = iota
)

var DB = map[SpritesheetId]*Spritesheet{
	PlayerSpritesheedId: NewSpritesheet(2, 3, 15, 26),
}
