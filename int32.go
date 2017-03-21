package atomic

import "sync/atomic"

// Int32 provides AddInt32, StoreInt32, LoadInt32 and CompareAndSwapInt32
// from the package sync/atomic as convenient methods on a type.
// An Int32 can be created as part of other structures. The zero value for an Int32 is 0.
// An Int32 must not be copied after first use.
type Int32 struct {
	i int32
}

// NewInt32 returns a new atomic int32 set inititially to value.
func NewInt32(value int32) *Int32 {
	return &Int32{i: value}
}

// Add is the atomic equivalent of:
//	*i32 += delta
//	return *i32
func (i32 *Int32) Add(delta int32) int32 {
	// generate nil checks and faults early.
	ref := &i32.i
	return atomic.AddInt32(ref, delta)
}

// Load is the atomic equivalent of:
//	return *i32
func (i32 *Int32) Load() int32 {
	// generate nil checks and faults early.
	ref := &i32.i
	return atomic.LoadInt32(ref)
}

// Store is the atomic equivalent of:
//	*i32 = value
func (i32 *Int32) Store(value int32) {
	// generate nil checks and faults early.
	ref := &i32.i
	atomic.StoreInt32(ref, value)
}

// CompareAndSwap executes the compare-and-swap operation on i32,
// which is the atomic equivalent of:
// 	if *i32 == old {
//		*i32 = new
//		return true
//	}
//	return false
func (i32 *Int32) CompareAndSwap(old, new int32) (swapped bool) {
	// generate nil checks and faults early.
	ref := &i32.i
	return atomic.CompareAndSwapInt32(ref, old, new)
}
