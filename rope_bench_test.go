package rope_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/zyedidia/rope"
)

var dummy int

func BenchmarkConstruction(b *testing.B) {
	rope.SplitLength = 4096 * 4
	rope.JoinLength = rope.SplitLength / 2

	benchmarks := []struct {
		size int
	}{
		{1000},
		{10000},
		{100000},
		{1000000},
		{10000000},
	}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("Size%d", bm.size), func(b *testing.B) {
			bytes := randbytes(bm.size)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				r := rope.New(bytes)
				dummy = r.Len()
			}
		})
	}
}

func BenchmarkSplice(b *testing.B) {
	rope.SplitLength = 4096 * 4
	rope.JoinLength = rope.SplitLength / 2

	benchmarks := []struct {
		size   int
		splice int
	}{
		{1000, 10},
		{10000, 10},
		{100000, 10},
		{1000000, 10},
		{10000000, 10},
	}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("Size%d:Splice%d", bm.size, bm.splice), func(b *testing.B) {
			bytes := randbytes(bm.size)
			r := rope.New(bytes)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				low := rand.Intn(bm.size)
				high := min(low+bm.splice, bm.size)
				slc := r.Slice(low, high)
				r.Remove(low, high)
				r.Insert(low, slc)
			}
		})
	}
}

func BenchmarkSpliceString(b *testing.B) {
	rope.SplitLength = 4096 * 4
	rope.JoinLength = rope.SplitLength / 2

	benchmarks := []struct {
		size   int
		splice int
	}{
		{1000, 10},
		{10000, 10},
		{100000, 10},
		{1000000, 10},
		{10000000, 10},
	}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("Size%d:Splice%d", bm.size, bm.splice), func(b *testing.B) {
			bytes := randbytes(bm.size)
			r := newBasicText(bytes)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				low := rand.Intn(bm.size)
				high := min(low+bm.splice, bm.size)
				slc := r.slice(low, high)
				r.remove(low, high)
				r.insert(low, slc)
			}
		})
	}
}
