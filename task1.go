package main

import (
	"fmt"
	"sort"
)

func task1() {
	fmt.Println("task1")

	fmt.Println("只出现一次的数字：", singleNumber([]int{4, 1, 2, 1, 2}))

	fmt.Println("回文数：", isPalindrome(121))

	fmt.Println("有效的括号：", isValid("()[]{}"))

	fmt.Println("最长公共前缀：", longestCommonPrefix([]string{"flower", "flow", "flight"}))

	fmt.Println("加一：", plusOne([]int{9, 9, 9}))

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	fmt.Println("删除排序数组中的重复项：", k, nums[:k])

	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merged := merge(intervals)
	fmt.Println("合并区间：", merged)

	nums2 := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums2, target)
	fmt.Println("两数之和：", result)

}

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 说明：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
/*
示例 1:
输入: [2,2,1]
输出: 1*/
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		// 异或操作
		// 相同的数异或结果为0，任何数与0异或结果为其本身
		// 因此成对出现的数最终会被抵消，剩下的就是只出现一次的数
		result ^= num
	}
	return result
}

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 例如，121 是回文，而 123 不是。
func isPalindrome(x int) bool {
	// 负数不是回文
	if x < 0 {
		return false
	}
	// 反转数字
	reverted := 0
	original := x
	for x > 0 {
		reverted = reverted*10 + x%10
		x /= 10
	}
	return original == reverted
}

// 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 1.左括号必须用相同类型的右括号闭合。
// 2.左括号必须以正确的顺序闭合。
// 3.每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
	stack := []rune{}
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix) && j < len(strs[i]); j++ {
			if prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
	}
	return prefix
}

// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
// 将大整数加 1，并返回结果的数字数组。
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

// 删除排序数组中的重复项
// 给定一个排序数组 nums ，删除重复元素，使每个元素只出现一次，返回新数组的长度。
// 元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
// 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
// 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
// 返回 k 。
// 判题标准:
// 系统会用下面的代码来测试你的题解:
// int[] nums = [...]; // 输入数组
// int[] expectedNums = [...]; // 长度正确的期望答案
// int k = removeDuplicates(nums); // 调用
// assert k == expectedNums.length;
//
//	for (int i = 0; i < k; i++) {
//	    assert nums[i] == expectedNums[i];
//	}
//
// 如果所有断言都通过，那么您的题解将被 通过。
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

// 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，且该数组需恰好覆盖输入中的所有区间 。
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	// 先按区间起点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		// 如果当前区间与上一个合并区间重叠，则合并
		if intervals[i][0] <= merged[len(merged)-1][1] {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], intervals[i][1])
		} else {
			// 否则，直接添加到结果中
			merged = append(merged, intervals[i])
		}
	}
	return merged
}

// 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
// 你可以按任意顺序返回答案。
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}
