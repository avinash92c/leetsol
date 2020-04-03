func subsearch(nums []int, target int, start, end, mid int) int {
	// println(start, end, mid)
	// println("Nums at Indexes")
	// println(nums[start], nums[end], nums[mid])
	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	} else if nums[mid] == target {
		return mid
	} else if start == end && nums[start] != target {
		return -1
	} else {
		// println("else block")
		locmid := (start + end) / 2
		if nums[locmid] == target {
			return locmid
		}
		// println("newstarts")
		// println(start, locmid, (start+locmid)/2)
		res := subsearch(nums, target, start, locmid, (start+locmid)/2)
		if res != -1 {
			return res
		}

		// println("newstarts-2")
		// println(locmid)
		// println(locmid+1, end, (locmid+1+end)/2)
		res = subsearch(nums, target, locmid+1, end, (locmid+1+end)/2)
		if res != -1 {
			return res
		}
		return -1
	}
}

func search(nums []int, target int) int { //FIND OUT HOW TO RETURN INDEX
	return subsearch(nums, target, 0, len(nums)-1, len(nums)/2)
}
