package algorithm

import (
	"fmt"
	"testing"
)

func TestMaxSubArraySum(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	output := MaxSubArraySum(input)
	fmt.Println(output) // 输出：6，因为 [4, -1, 2, 1] 的和最大
}
