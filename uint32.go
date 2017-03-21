package atomic

import "sync/atomic"

// Uint32 provides AddUint32, StoreUint32, LoadUint32 and CompareAndSwapUint32
// from the package sync/atomic as convenient methods on a type.
// An Uint32 can be created as part of other structures. The zero value for an Uint32 is 0.
// An Uint32 must not be copied after first use.
type Uint32 struct {
	i uint32
}

// NewUint32 returns a new atomic uint32 set inititially to value.
func NewUint32(value uint32) *Uint32 {
	return &Uint32{i: value}
}

// Add is the atomic equivalent of:
//	*u32 += delta
//	return *u32
func (u32 *Uint32) Add(delta uint32) uint32 {
	// generate nil checks and faults early.
	ref := &u32.i
	return atomic.AddUint32(ref, delta)
}

// Load is the atomic equivalent of:
//	return *u32
func (u32 *Uint32) Load() uint32 {
	// generate nil checks and faults early.
	ref := &u32.i
	return atomic.LoadUint32(ref)
}

// Store is the atomic equivalent of:
//	*u32 = value
func (u32 *Uint32) Store(value uint32) {
	// generate nil checks and faults early.
	ref := &u32.i
	atomic.StoreUint32(ref, value)
}

// CompareAndSwap executes the compare-and-swap operation on u32,
// which is the atomic equivalent of:
// 	if *u32 == old {
//		*u32 = new
//		return true
//	}
//	return false
func (u32 *Uint32) CompareAndSwap(old, new uint32) (swapped bool) {
	// generate nil checks and faults early.
	ref := &u32.i
	return atomic.CompareAndSwapUint32(ref, old, new)
}
