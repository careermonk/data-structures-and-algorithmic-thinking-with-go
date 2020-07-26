// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
//

package main

import (
	"crypto/rand"
	"fmt"
	"hash/fnv"
	"math"
)

type BloomFilter struct {
	Capacity          int   // n
	FalsePositiveRate int   // p
	NumHashes         int   // k, Number of hash functions
	BitSize           int64 // m, Size of the bloom filter
	numBuckets        int64
	state             []uint32
}

// You can simulate additonal hash functions with two hash functions.
// Futhermore we can simulate 2x32-bit hash functions with one 64-bit hash.
func hashFNV1a(input []byte) (uint32, uint32) {
	hash := fnv.New64a()
	hash.Write(input)
	value64 := hash.Sum64()
	return uint32(value64 & 0xFFFFFFFF), uint32(value64 >> 32)
}

func (set *BloomFilter) setBit(index int64) {
	bucket := (index / 32) % set.numBuckets
	offset := index % 32
	set.state[bucket] = set.state[bucket] | (1 << uint(offset))
}

func (set *BloomFilter) testBit(index int64) int {
	bucket := (index / 32) % set.numBuckets
	offset := index % 32
	if set.state[bucket]&(1<<uint(offset)) != 0 {
		return 1
	} else {
		return 0
	}
}

// Adds a piece of arbitrary data to the set
func (set *BloomFilter) Put(input []byte) {
	hashA, hashB := hashFNV1a(input)
	for i := 0; i < set.NumHashes; i++ {
		index := int64((hashA + hashB*uint32(i))) % set.BitSize
		set.setBit(index)
	}
}

// Checks the set for a piece of arbitrary data
func (set *BloomFilter) MightContains(input []byte) bool {
	hashA, hashB := hashFNV1a(input)
	for i := 0; i < set.NumHashes; i++ {
		index := int64((hashA + hashB*uint32(i))) % set.BitSize
		if set.testBit(index) != 1 {
			return false
		}
	}
	return true
}

// Returns a new Bloom filter. Parameters are the expected number of elements
// in the set and the desired false positive probability. Optimal size and
// number of hashes are calculated based on these numbers.
//
// p = false positive rate of the form 1/p, powers of two preferred
// optimal number of hashes k = (m/n)ln(2)
func NewBloomFilter(capacity, probability int) *BloomFilter {
	bitSize := int64(math.Abs(math.Ceil(float64(capacity) *
		math.Log2(math.E) * math.Log2(1/float64(probability)))))
	numHashes := int(math.Floor(float64((bitSize / int64(capacity))) * math.Log(2)))
	numBuckets := bitSize / 32
	return &BloomFilter{
		Capacity:          capacity,
		FalsePositiveRate: probability,
		NumHashes:         numHashes,
		BitSize:           bitSize,
		numBuckets:        numBuckets,
		state:             make([]uint32, uint(numBuckets))}
}

func main() {
	var trials = 1000
	var rate = 1000
	var counter = 0

	// False Nagatives
	set := NewBloomFilter(trials, rate)

	// Add random data to set
	for i := 0; i < trials; i++ {
		c := 8
		b := make([]byte, c)
		_, err := rand.Read(b)
		if err != nil {
			fmt.Println("Could not generate random string")
			return
		}
		set.Put(b)
		if set.MightContains(b) {
			counter++
		}
	}

	if trials != counter {
		fmt.Println("Lost some random set entries!")
		return
	}

	// FalsePositives
	set = NewBloomFilter(trials, rate)

	// Add random data to set
	for i := 0; i < trials; i++ {
		c := 8
		b := make([]byte, c)
		_, err := rand.Read(b)
		if err != nil {
			fmt.Println("Could not generate random string")
			return
		}
		set.Put(b)
	}

	// Check for different random data
	for i := 0; i < trials; i++ {
		c := 8
		b := make([]byte, c)
		_, err := rand.Read(b)
		if err != nil {
			fmt.Println("Could not generate random string")
			return
		}
		if set.MightContains(b) {
			counter++
		}
	}

	expected := float64(1) / float64(rate)
	actual := float64(counter) / float64(trials)

	fmt.Println("expected error rate: ", expected)
	fmt.Println("actual error rate: ", actual)
}
