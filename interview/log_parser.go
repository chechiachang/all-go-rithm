package interview

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ParseLog(s1, s2 string, order int) string {
	slice1 := strings.Split(s1, ",")
	slice2 := strings.Split(s2, ",")

	m := map[int]string{}
	keys := []int{}

	for _, s := range slice1 {
		keyValue := strings.Split(s, ":")
		if k, err := strconv.Atoi(keyValue[0]); err != nil {
			fmt.Println(err)
			return ""
		} else {
			m[k] = s
			keys = append(keys, k)
		}
	}

	// TODO
	for _, s := range slice2 {
		keyValue := strings.Split(s, ":")
		if k, err := strconv.Atoi(keyValue[0]); err != nil {
			fmt.Println(err)
			return ""
		} else {
			m[k] = s
			keys = append(keys, k)
		}
	}

	sort.Ints(keys)

	return m[keys[order]]
}
