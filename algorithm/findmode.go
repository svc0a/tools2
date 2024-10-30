package algorithm

//计算一个整数数组中所有元素的众数。众数是数组中出现次数最多的元素，如果有多个数出现相同的次数，则返回任意一个众数即可。
//
//任务描述
//编写一个函数 findMode(nums []int) int，接受一个整数数组 nums，并返回其中的众数。
//
//示例
//go
//Copy code
//input := []int{1, 2, 3, 2, 2, 3, 3, 3}
//output := findMode(input)
//fmt.Println(output) // 输出：3
//提示
//使用一个 map 来统计每个数的出现次数。
//遍历 map 找出出现次数最多的数。
//期望的代码实现
//试试编写 findMode 函数，或者如果遇到问题，我可以进一步给你提供提示！

func FindMode(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	result := m[0]
	for k, v := range m {
		if v > result {
			result = k
		}
	}
	return result
}
