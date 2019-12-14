package containers

import "github.com/emirpasic/gods/utils"

type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}

func GetStoredValue(container Container, comprartor utils.Comparator) []interface{} {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	utils.Sort(values, comparator)
	return values
}
