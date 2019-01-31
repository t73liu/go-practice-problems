package array

func intersect(nums1 []int, nums2 []int) []int {
	count1 := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		val, _ := count1[nums1[i]]
		count1[nums1[i]] = val + 1
	}

	intersect := make([]int, 0)

	for i := 0; i < len(nums2); i++ {
		if len(count1) == 0 {
			break
		}
		val, ok := count1[nums2[i]]
		if ok {
			intersect = append(intersect, nums2[i])
			if val == 1 {
				delete(count1, nums2[i])
			} else {
				count1[nums2[i]] = val - 1
			}
		}
	}

	return intersect
}
