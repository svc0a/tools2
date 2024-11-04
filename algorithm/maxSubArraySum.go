package algorithm

// 题目：寻找数组的连续子数组最大和
// 编写一个函数 maxSubArraySum(nums []int) int，该函数接收一个整数数组 nums，返回连续子数组的最大和。这个问题也被称为 最大子序和问题（Maximum Subarray Problem）。
//
// input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
// output := maxSubArraySum(input)
// fmt.Println(output) // 输出：6，因为 [4, -1, 2, 1] 的和最大
// 提示
// 尝试使用 Kadane’s Algorithm，这个算法的时间复杂度是 𝑂(𝑛)。
// 通过逐个累加数组的值来维护一个当前子数组的和，当和变为负数时，重置它为零。
func MaxSubArraySum(nums []int) int {
	sumSoFar := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sumSoFar > 0 {
			sumSoFar += nums[i]
		} else {
			sumSoFar = nums[i]
		}
		sum = max(sumSoFar, sum)
	}
	return sum
}

//算法原理
//Kadane的算法通过一次遍历数组来实现，它维护两个变量：maxEndingHere和maxSoFar。maxEndingHere存储到当前元素为止的最大子数组和，而maxSoFar则记录遍历到目前为止所有子数组的最大和。
//对于数组中的每一个元素，Kadane的算法都会计算两个选择：
//1. 将当前元素加入到前一个元素的最大子数组中（如果前一个元素的最大子数组和为正）。
//2. 从当前元素开始一个新的子数组（如果前一个元素的最大子数组和为负或者开始一个新的子数组更有利）。
//算法的每一步都会更新maxEndingHere为这两个选择中的较大值，并且可能更新maxSoFar为当前的maxEndingHere和已知的maxSoFar中的较大值。
