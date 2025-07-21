package main

// here we need to find the element in roated array=
// array start with left and right
// then we check with mid value, which side have the small value
// then after we go through for that
// when we reach the min value so that time left and right will same index

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// If the middle element is greater than the rightmost element,
		// the minimum is in the right half.
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			// Otherwise, the minimum is in the left half (including mid).
			right = mid
		}
	}

	// At the end of the loop, left == right, pointing to the minimum.
	return nums[left]
}
