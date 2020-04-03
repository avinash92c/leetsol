func intersect(nums1 []int, nums2 []int) []int {
    var res []int
  mapsize:=0
	if len(nums1)>len(nums2){
		mapsize = len(nums1)
	}else{
		mapsize = len(nums2)
	}
	nummap := make(map[int]int, mapsize)
	for _, num1 := range nums2 {
		_, ok := nummap[num1]
		if ok {
			nummap[num1]++
		} else {
			nummap[num1] = 1
		}
	}

	for _, num2 := range nums1 {
		if counter, ok := nummap[num2]; counter>0 && ok {
			res = append(res, num2)
			nummap[num2]--
		}
	}
    return res
}
