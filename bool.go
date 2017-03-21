package atomic

import "sync/atomic"

// Bool is an atomic boolean.
// It can be flipped from known states to newly expected states and queries.
// A Bool can be created as part of other structures. The zero value for a Bool is false.
// A Bool must not be copied after first use.
type Bool struct {
	i uint64
}

// NewBool returns a new atomic boolean set inititially to value.
func NewBool(value bool) *Bool {
	if value == true {
		return &Bool{i: 1}
	}
	return &Bool{}
}

// Value returns the current boolean value.
func (b *Bool) Value() bool {
	// generate nil checks and faults early.
	ref := &b.i
	return atomic.LoadUint64(ref) == 1
}

// CompareAndSwap executes the compare-and-swap operation on bool,
// which is the atomic equivalent of:
// 	if *b == old {
//		*b = new
//		return true
//	}
//	return false
func (b *Bool) CompareAndSwap(old, new bool) (swapped bool) {
	// generate nil checks and faults early.
	ref := &b.i

	switch {
	case old == false && new == true:
		return atomic.CompareAndSwapUint64(ref, 0, 1)
	case old == true && new == false:
		return atomic.CompareAndSwapUint64(ref, 1, 0)
	default:
		// this case is nonsense, since a swap will never happen.
		// new == old
		return false
	}
}
