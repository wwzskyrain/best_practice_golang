package main

// 2789. Largest Element in an Array after Merge Operations
func maxArrayValue(nums []int) int64 {
	l := len(nums)
	sum := int64(nums[l-1])
	for i := l - 2; i >= 0; i-- {
		n := int64(nums[i])
		if sum >= n {
			sum += n
		} else {
			sum = n
		}
	}
	return sum
}
