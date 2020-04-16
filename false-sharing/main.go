/*
@Time : 2020/4/16 2:08 PM
*/

package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	myatomic := &Pad{}
	myatomic.IncreaseAllElements()
	fmt.Printf("%d\n", myatomic.A)
}

type MyAtomic interface {
	IncreaseAllElements()
	IncreaseA()
	IncreaseB()
}
type NoPad struct {
	A uint64
	B uint64
	C uint64
}

func (p *NoPad) IncreaseA() {
	atomic.AddUint64(&p.A, 1)
}

func (p *NoPad) IncreaseB() {
	atomic.AddUint64(&p.B, 1)
}

func (p *NoPad) IncreaseAllElements() {
	atomic.AddUint64(&p.A, 1)
	atomic.AddUint64(&p.B, 1)
	atomic.AddUint64(&p.C, 1)
}

type Pad struct {
	A uint64
	_ [8]uint64
	B uint64
	_ [8]uint64
	C uint64
	_ [8]uint64
}

func (p *Pad) IncreaseA() {
	atomic.AddUint64(&p.A, 1)
}

func (p *Pad) IncreaseB() {
	atomic.AddUint64(&p.B, 1)
}

func (p *Pad) IncreaseAllElements() {
	atomic.AddUint64(&p.A, 1)
	atomic.AddUint64(&p.B, 1)
	atomic.AddUint64(&p.C, 1)
}
