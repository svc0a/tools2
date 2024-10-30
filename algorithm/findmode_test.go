package algorithm

import (
	"fmt"
	"testing"
)

func TestFindMode(t *testing.T) {
	input := []int{1, 2, 3, 2, 2, 3, 3, 3}
	output := FindMode(input)
	fmt.Println(output) // 输出：3
}
