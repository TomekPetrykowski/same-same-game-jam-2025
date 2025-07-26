package scene

import (
	"game/entities"
	"sort"
)

func GetSortedObjects(objects *map[entities.SceneObjectId][]entities.GameObject) *[]entities.GameObject {
	var allObjectsLen = 0
	for _, slice := range *objects {
		allObjectsLen += len(slice)
	}

	var allObjects = make([]entities.GameObject, 0, allObjectsLen)

	for _, slice := range *objects {
		allObjects = append(allObjects, slice...)
	}

	sort.Slice(allObjects, func(i, j int) bool {
		iY := allObjects[i].GetCollider().GetYForDrawing()
		jY := allObjects[j].GetCollider().GetYForDrawing()
		return iY < jY
	})

	return &allObjects
}
