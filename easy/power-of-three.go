func isPowerOfThree(n int) bool {
	if n % 2==0 || n<0 {
		return false
	}
    
	for {
		mod := n % 3
		n = n / 3
		// print(n)
		// print(" - ")
		// println(mod)
		if mod > 1 || (mod == 1 && n == 1) || (mod==1 && n>1){
			return false
		} else if mod == 1 && n==0 {
			return true
		}
	}
	// return true
}
