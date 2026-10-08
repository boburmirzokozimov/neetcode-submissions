const bucketsCount = 1000

type MyHashMap struct {
	buckets [][]Pair
}
type Pair struct {
	key int
	value int
}

func Constructor() MyHashMap {
    return MyHashMap{
		buckets: make([][]Pair, bucketsCount),
	}
}

func (this *MyHashMap) Put(key int, value int) {
    h := hash(key)
	for i := range this.buckets[h] {
		if this.buckets[h][i].key == key {
			this.buckets[h][i].value = value
			return
		}
	}
	p := Pair{
		key: key,
		value: value,
	}
	this.buckets[h] = append(this.buckets[h], p)
}

func (this *MyHashMap) Get(key int) int {
    h := hash(key)
    for _, v := range this.buckets[h] {
		if v.key == key {
			return v.value
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
    h := hash(key)
    for k, v := range this.buckets[h] {
		if v.key == key {
			this.buckets[h] = append(
				this.buckets[h][:k],
				this.buckets[h][k+1:]...
			)
			return
		}
	}
}

func hash(k int) int {
	return k % bucketsCount
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */