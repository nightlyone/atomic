package atomic

import "sync/atomic"

// Uint64 provides AddUint64, StoreUint64, LoadUint64 and CompareAndSwapUint64
// from the package sync/atomic as convenient methods on a type.
// An Uint64 can be created as part of other structures. The zero value for an Uint64 is 0.
// An Uint64 must not be copied after first use.
type Uint64 struct {
	i uint64
}

// NewUint64 returns a new atomic uint64 set inititially to value.
func NewUint64(value uint64) *Uint64 {
	return &Uint64{i: value}
}

// Add is the atomic equivalent of:
//	*u64 += delta
//	return *u64
func (u64 *Uint64) Add(delta uint64) uint64 {
	// generate nil checks and faults early.
	ref := &u64.i
	return atomic.AddUint64(ref, delta)
}

// Load is the atomic equivalent of:
//	return *u64
func (u64 *Uint64) Load() uint64 {
	// generate nil checks and faults early.
	ref := &u64.i
	return atomic.LoadUint64(ref)
}

// Store is the atomic equivalent of:
//	*u64 = value
func (u64 *Uint64) Store(value uint64) {
	// generate nil checks and faults early.
	ref := &u64.i
	atomic.StoreUint64(ref, value)
}

// CompareAndSwap executes the compare-and-swap operation on u64,
// which is the atomic equivalent of:
// 	if *u64 == old {
//		*u64 = new
//		return true
//	}
//	return false
func (u64 *Uint64) CompareAndSwap(old, new uint64) (swapped bool) {
	// generate nil checks and faults early.
	ref := &u64.i
	return atomic.CompareAndSwapUint64(ref, old, new)
}
