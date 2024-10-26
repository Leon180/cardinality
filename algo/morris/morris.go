package morris

import (
	"math/rand"
)

// MorrisCounterTest struct holds the MorrisCounter in a slice
type MorrisCounterTest struct {
	morrisCounter MorrisCounter
}

func NewMorrisCounterTest(
	counts int,
) *MorrisCounterTest {
	mc := NewMorrisCounter(counts)
	return &MorrisCounterTest{
		morrisCounter: mc,
	}
}

func (mct *MorrisCounterTest) Run() {
	if mct == nil {
		return
	}
	mct.morrisCounter.run()
}

func (mct *MorrisCounterTest) GetEstimateCounts() []int {
	if mct == nil {
		return nil
	}
	return mct.morrisCounter.getEstimateCounts()
}

type MorrisCounter []int

func NewMorrisCounter(counts int) MorrisCounter {
	record := make([]int, counts)
	// init the first value as 1
	if counts > 0 {
		record[0] = 1
	}
	return MorrisCounter(record)
}

func (mc MorrisCounter) run() {
	for i := 1; i < len(mc); i++ {
		mc.event(i)
	}
}

func (mc MorrisCounter) event(i int) {
	randVal := rand.Float64()
	if randVal < (1.0 / float64(int(1)<<mc[i-1])) {
		mc[i] = mc[i-1] + 1
	} else {
		mc[i] = mc[i-1]
	}
}

func (mc MorrisCounter) getEstimateCounts() []int {
	counts := make([]int, len(mc))
	for i, count := range mc {
		counts[i] = mc.estimateCount(count)
	}
	return counts
}

func (mc *MorrisCounter) estimateCount(count int) int {
	return (1 << count) - 1 // 2^count - 1
}
