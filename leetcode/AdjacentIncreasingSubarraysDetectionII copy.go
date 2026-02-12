package leetcode

func maxIncreasingSubarrays(nums []int) int {

	i := 1
	leftLen := 0
	k := 1
	for i < len(nums) {
		curLen := 1
		for i < len(nums) {
			if nums[i] > nums[i-1] {
				curLen++
			} else {
				break
			}
			i++
		}
		i++
		minLen := min(leftLen, curLen)
		if minLen > k {
			k = minLen
		}
		if curLen/2 > k {
			k = curLen / 2
		}
		leftLen = curLen
	}
	return k
}
