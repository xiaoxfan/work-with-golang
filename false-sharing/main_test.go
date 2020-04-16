/*
@Time : 2020/4/16 2:12 PM
*/
package main

import (
	"sync"
	"testing"
)

func testAtomicIncrease(myatomic MyAtomic) {
	paraNum := 1000
	addTimes := 1000
	var wg sync.WaitGroup
	wg.Add(paraNum)
	for i := 0; i < paraNum; i++ {
		go func() {
			for j := 0; j < addTimes; j++ {
				myatomic.IncreaseAllElements()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkNoPad(b *testing.B) {
	myatomic := &NoPad{}
	b.ResetTimer()
	testAtomicIncrease(myatomic)
}

func BenchmarkPad(b *testing.B) {
	myatomic := &Pad{}
	b.ResetTimer()
	testAtomicIncrease(myatomic)
}
//BenchmarkNoPad-4        1000000000               0.0580 ns/op
//BenchmarkPad-4          1000000000               0.0203 ns/op