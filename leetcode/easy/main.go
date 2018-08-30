package easy

import (
	"fmt"
)

func main() {
	x := 14
	for i := 0; i < 10; i++ {
		x = x >> 1
		fmt.Println(x)
	}
}
