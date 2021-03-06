// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                  priveda/fixed64/[sub_int64.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// SubInt64 subtracts one or more 64-bit integers from a fixed-point
// number and returns the result. The original number is not changed.
func (n Fixed64) SubInt64(nums ...int64) Fixed64 {
	for _, num := range nums {
		n.i64 -= num * 1e4
	}
	return n
}

// end
