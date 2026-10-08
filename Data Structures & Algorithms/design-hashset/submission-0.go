type MyHashSet struct {
	data map[int]bool
}

func Constructor() MyHashSet {
    return MyHashSet{
		data: make(map[int]bool),
	}
}

func (this *MyHashSet) Add(key int) {
    this.data[key] = true
}

func (this *MyHashSet) Remove(key int) {
    delete(this.data, key)
}

func (this *MyHashSet) Contains(key int) bool {
	_, exists := this.data[key]
	return exists
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 