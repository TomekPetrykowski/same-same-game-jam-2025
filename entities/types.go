package entities

type Scene interface {
	GetObjects() *map[string][]GameObject
	AddObject(string, GameObject)
}
