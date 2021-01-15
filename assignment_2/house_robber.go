package assignment_2

// [1, 2, 3, 1]
func HouseRob(nums []int) int {
	evenSum := 0
	oddSum := 0
	for i := 0; i < len(nums)-1; i++ {
		if i%2 == 0 {
			evenSum += nums[i]
		} else {
			oddSum += nums[i]
		}
	}
	if evenSum > oddSum {
		return evenSum
	}
	return oddSum
}
