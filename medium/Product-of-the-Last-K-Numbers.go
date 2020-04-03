type ProductOfNumbers struct {
    nums []int
}


func Constructor() ProductOfNumbers {
    nums := make([]int,0)
    return ProductOfNumbers{nums:nums}
}


func (this *ProductOfNumbers) Add(num int)  {
    this.nums = append(this.nums,num)
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    start := len(this.nums) - k
	sum := 1
    
	for start < len(this.nums) {
		sum *= this.nums[start]
		start++
	}
	return sum
}


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
