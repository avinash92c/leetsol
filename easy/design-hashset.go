type MyHashSet struct {
    Mhs map[int]int
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
    return MyHashSet{Mhs: make(map[int]int),}
}


func (this *MyHashSet) Add(key int)  {
    this.Mhs[key]=1
}


func (this *MyHashSet) Remove(key int)  {
    delete(this.Mhs,key)
    //this.Mhs.Delete(key)
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
//    return this.Mhs.Contains(key)
    _,ok := this.Mhs[key]
    return ok
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
