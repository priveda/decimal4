// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-15 16:21:27 2B239F                    priveda/fixed64/[module.go]
// -----------------------------------------------------------------------------

package fixed64

// # Fixed64 Type:
//   Fixed64 struct
//
// # Fixed64 Factories:
//   Fixed64Of(value interface{}) Fixed64
//   Fixed64Raw(raw int64) Fixed64
//   ParseFixed64(s string) (Fixed64, error)
//
// # String Output:
//   (n Fixed64) GoString() string
//   (n Fixed64) Fmt(decimalPlaces int) string
//   (n Fixed64) InWordsEN(fmt string) string
//   (n Fixed64) String() string
//
// # Division:
//   (n Fixed64) Div(nums ...Fixed64) Fixed64
//   (n Fixed64) DivFloat(nums ...float64) Fixed64
//   (n Fixed64) DivInt(nums ...int) Fixed64
//   (n Fixed64) DivInt64(nums ...int64) Fixed64
//
// # Multiplication:
//   (n Fixed64) Mul(nums ...Fixed64) Fixed64
//   (n Fixed64) MulFloat(nums ...float64) Fixed64
//   (n Fixed64) MulInt(nums ...int) Fixed64
//   (n Fixed64) MulInt64(nums ...int64) Fixed64
//
// # Addition:
//   (n Fixed64) Add(nums ...Fixed64) Fixed64
//   (n Fixed64) AddFloat(nums ...float64) Fixed64
//   (n Fixed64) AddInt(nums ...int) Fixed64
//   (n Fixed64) AddInt64(nums ...int64) Fixed64
//
// # Subtraction:
//   (n Fixed64) Sub(nums ...Fixed64) Fixed64
//   (n Fixed64) SubFloat(nums ...float64) Fixed64
//   (n Fixed64) SubInt(nums ...int) Fixed64
//   (n Fixed64) SubInt64(nums ...int64) Fixed64
//
// # Information:
//   (n Fixed64) Float64() float64
//   (n Fixed64) Int() int
//   (n Fixed64) Int64() int64
//   (n Fixed64) IsEqual(value interface{}) bool
//   (n Fixed64) IsGreater(value interface{}) bool
//   (n Fixed64) IsGreaterOrEqual(value interface{}) bool
//   (n Fixed64) IsLesser(value interface{}) bool
//   (n Fixed64) IsLesserOrEqual(value interface{}) bool
//   (n Fixed64) IsNegative() bool
//   (n Fixed64) IsZero() bool
//   (n Fixed64) Overflow() int
//   (n Fixed64) Raw() int64
//
// # JSON:
//   (n Fixed64) MarshalJSON() ([]byte, error)
//   (n *Fixed64) UnmarshalJSON(data []byte) error
//
// # Helper Function
//   fixed64Overflow(isNegative bool, a ...interface{}) Fixed64

// -----------------------------------------------------------------------------
// # Function Proxy Variables (for mocking)

type thisMod struct {
	Error func(args ...interface{}) error
}

var mod = thisMod{
	Error: logError,
}

// ModReset restores all mocked functions to the original standard functions.
func (ob *thisMod) Reset() {
	ob.Error = logError
}

// logError __
func logError(args ...interface{}) error {
	return nil
}

//end
