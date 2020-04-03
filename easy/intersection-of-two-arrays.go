func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	nummap := make(map[int]bool,len(nums1))
	for _,num1 := range nums1{
		nummap[num1]= true
	}

	for _,num2 := range nums2{
		if _,ok:=nummap[num2]; ok && !slicecontains(res,num2) {
			res = append(res,num2)
		}
	}
	return res
}

func slicecontains(input []int,value int)bool{
	res:= false
 for _, val := range input{
	 if val == value{
		res = true
		break;
	 }

 }
 return res
}
