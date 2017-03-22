// Package atomic implements atomic operations as methods on types so they can be used more easily.
package atomic

// BUG(nightlyone): The implementation just wraps sync/atomic from the Go standard library,
// so it suffers the same implementation limitations.
