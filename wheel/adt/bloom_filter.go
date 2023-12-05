package adt

type BloomFilter[T any] struct {
	bitMap         *BitMap
	bitCounts      uint64
	elemCounts     uint64
	hashFuncCounts uint8
	hashFunc       []func(T) uint64
}

func MakeBloomFilter[T any](assumeElemCounts uint64, hashFunc ...func(T) uint64) *BloomFilter[T] {
	bf := &BloomFilter[T]{}
	return bf
}

func (bf *BloomFilter[T]) make(elemCounts uint64, hashFunc ...func(T) uint64) {
	bf.elemCounts = bf.uint64CeilBin(elemCounts)
}

func (bf *BloomFilter[T]) Insert(t T) {
	bf.insert(t)
}

func (bf *BloomFilter[T]) insert(t T) {

}

func (bf *BloomFilter[T]) Query(t T) bool {
	return bf.query(t)
}

func (bf *BloomFilter[T]) query(t T) bool {
	return false
}

func (bf *BloomFilter[T]) uint64CeilBin(x uint64) uint64 {
	return x
}
