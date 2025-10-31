package piscine

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if i > 2 {
			break
		}
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums[2]
}
