/*
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
*/
package leetcode

// map
func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for index, val := range nums {
		if i, ok := hashTable[target-val]; ok {
			return []int{i, index}
		}
		hashTable[val] = index
	}
	return nil
}

// 两遍循环
func twoSum1(nums []int, target int) []int {
	for index1, val1 := range nums {
		for index2, val2 := range nums {
			if target == val1+val2 {
				return []int{index1, index2}
			}
		}
	}
	return nil
}
