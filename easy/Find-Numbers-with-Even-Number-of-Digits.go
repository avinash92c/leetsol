func findNumbers(nums []int) int {
    evenNums := 0
    for _, v := range nums{
        if len(strconv.Itoa(v)) % 2 == 0{
            evenNums++
        }
    }
    
    return evenNums
}
