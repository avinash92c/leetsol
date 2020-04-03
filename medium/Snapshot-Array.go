package main

type SnapshotArray struct {
	snapshot []map[int]int
	actual   map[int]int
	length   int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		actual:   make(map[int]int),
		snapshot: make([]map[int]int, 0),
		length:   length,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.actual[index] = val
}

func (this *SnapshotArray) Snap() int {
	if len(this.actual) == 0 {
		this.snapshot = append(this.snapshot, nil)
	} else {
		// tmp := make([]int, this.length)
		tmp := make(map[int]int)
		for k, v := range this.actual {
			tmp[k] = v
		}
		this.snapshot = append(this.snapshot, tmp)
	}
	// this.actual = make(map[int]int)
	return len(this.snapshot) - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	snap := this.snapshot[snap_id]
	if snap != nil {
		val, ok := snap[index]
		if len(snap) == 0 || !ok {
			return 0
		}
		return val
	}
	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
