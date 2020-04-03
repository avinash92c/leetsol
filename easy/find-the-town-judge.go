func findJudge(N int, trust [][]int) int {
    m := make([][]int, N)
    for i := 0; i < N; i++ {
        m[i] = make([]int, N)
    }
    
    for _,val:= range trust{
        l,r:= val[0],val[1]
        m[l-1][r-1]++
    }
    
    sideidx:= -1
    
    for i:=0;i<len(m);i++{
        allzero:= true
        for _,val:= range m[i]{
            if val==1{
                allzero = false
                break
            }
        }
        
        if allzero{
            sideidx = i
            break
        }
    }
    
    if sideidx==-1{
        return -1
    }
    
    
    sumval := 0
    for i:=0;i<len(m);i++{
        val := m[i][sideidx]
        sumval+=val
    }
    
    if sumval==N-1{
        return sideidx+1
    }
    
    return -1
}
