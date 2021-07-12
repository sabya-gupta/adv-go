package adv

func BinSearch(array []int, target int, lowIdx int, hightIdx int) int {
	if hightIdx < lowIdx {
		return -1
	}

	mid := int((hightIdx + lowIdx) / 2)

	if (array[mid]) == target {
		return mid
	}
	if mid > target {
		return BinSearch(array, target, lowIdx, mid)
	}

	if mid < target {
		return BinSearch(array, target, mid, hightIdx)
	}
	return -1
}
