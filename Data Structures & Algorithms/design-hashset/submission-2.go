const bucketsCount = 1000

type MyHashSet struct {
	buckets [][]int
}

func Constructor() MyHashSet {
    return MyHashSet{
		buckets: make([][]int, bucketsCount),
	}
}

func (this *MyHashSet) Add(key int) {
    h := hash(key)
	for _, v := range this.buckets[h] {
		if v == key {
			return
		}
	}
	this.buckets[h] = append(this.buckets[h], key)
}

func (this *MyHashSet) Remove(key int) {
    h := hash(key)
	
	for i, v := range this.buckets[h] {
		if v == key {
			this.buckets[h] = append(
				this.buckets[h][:i], 
				this.buckets[h][i+1:]...
				)
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
    h := hash(key)
	for _, v := range this.buckets[h] {
		if v == key {
			return true
		}
	}
	return false
}

func hash(key int) int {
    return key % bucketsCount
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 