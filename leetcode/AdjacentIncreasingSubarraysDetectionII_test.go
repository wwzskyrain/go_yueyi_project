package leetcode

import (
	"testing"
)

func Test2MaxIncreasingSubarrays(t *testing.T) {

	nums1 := []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}
	nums2 := []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}
	expects := []int{3, 2}
	numsArr := [][]int{nums1, nums2}
	for i, nums := range numsArr {
		expected := expects[i]
		result := maxIncreasingSubarrays(nums)
		if result != expected {
			t.Errorf("maxIncreasingSubarrays(%v) = %d; want %d", nums, result, expected)
		}
	}
}
