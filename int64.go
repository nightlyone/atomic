package atomic

import "sync/atomic"

// Int64 provides AddInt64, StoreInt64, LoadInt64 and CompareAndSwapInt64
// from the package sync/atomic as convenient methods on a type.
// An Int64 can be created as part of other structures. The zero value for an Int64 is 0.
// An Int64 must not be copied after first use.
type Int64 struct {
	i int64
}

// NewInt64 returns a new atomic int64 set inititially to value.
func NewInt64(value int64) *Int64 {
	return &Int64{i: value}
}

// Add is the atomic equivalent of:
//	*i64 += delta
//	return *i64
func (i64 *Int64) Add(delta int64) int64 {
	// generate nil checks and faults early.
	ref := &i64.i
	return atomic.AddInt64(ref, delta)
}

// Load is the atomic equivalent of:
//	return *i64
func (i64 *Int64) Load() int64 {
	// generate nil checks and faults early.
	ref := &i64.i
	return atomic.LoadInt64(ref)
}

// Store is the atomic equivalent of:
//	*i64 = value
func (i64 *Int64) Store(value int64) {
	// generate nil checks and faults early.
	ref := &i64.i
	atomic.StoreInt64(ref, value)
}

// CompareAndSwap executes the compare-and-swap operation on i64,
// which is the atomic equivalent of:
// 	if *i64 == old {
//		*i64 = new
//		return true
//	}
//	return false
func (i64 *Int64) CompareAndSwap(old, new int64) (swapped bool) {
	// generate nil checks and faults early.
	ref := &i64.i
	return atomic.CompareAndSwapInt64(ref, old, new)
}
