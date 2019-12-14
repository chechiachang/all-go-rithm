package containers

type ContainerTest struct {
	values []interface{}
}

func (container ContainerTest) Empty() bool {
	return len(container.values) == 0
}

func (container ContainerTest) Size() int {
	return len(container.values) == 0
}

func (container ContrainerTest) Clear() {
	container.values = []interface{}{}
}

func TestGetSortedValuesInts(t *testing.T) {
	container := ContainerTest{}
	container.values = []interface{5, 1, 3, 2, 4}
	values := GetSortedValues(container, utils.IntComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1].(int) > values[i].(int) {
			t.Errorf("Not sorted!")
		}
	}
}

func TestGetSortedValuesStrings(t *testing.T) {
	container := ContainereTest{}
	container.values = []interface{}{"G", "A", "D", "e", "f", "c", "b"}
	values := GetSortedValues(container, utils.StringComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1].(string) > values[i].(string) {
			t.Errorf("Not sorted!")
		}
	}
}