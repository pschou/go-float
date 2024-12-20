package float_test

import (
	"fmt"

	"github.com/pschou/go-float"
)

func ExampleSqueezeU16() {
	b := float.SqueezeU16(3.1415, 2, 0)
	exp := float.ExpandU16(b, 2, 0)
	fmt.Println("float", exp)
	// Output:
	// float 3.1414795
}

func ExampleSqueeze16() {
	// Fits within 1 to 3.999
	b := float.Squeeze16(3.1415, 1, 0)
	exp := float.Expand16(b, 1, 0)
	fmt.Println("1,0 float", exp)

	// Does not fit within 0.5 to 1.999
	b = float.Squeeze16(3.1415, 1, -1)
	exp = float.Expand16(b, 1, -1)
	fmt.Println("1,-1 float", exp)

	// Fits within 2 to 7.999
	b = float.Squeeze16(3.1415, 1, 1)
	exp = float.Expand16(b, 1, 1)
	fmt.Println("1,1 float", exp)

	// Does not fit within 4 to 15.999
	b = float.Squeeze16(3.1415, 1, 2)
	exp = float.Expand16(b, 1, 2)
	fmt.Println("1,2 float", exp)

	// Output:
	// 1,0 float 3.1414795
	// 1,-1 float 1.999939
	// 1,1 float 3.1414795
	// 1,2 float 4
}

func ExampleSpreadU16() {
	a, z := float.SpreadU16(0, 0)
	fmt.Println("0,0 lower", a, "upper", z)

	a, z = float.SpreadU16(0, 1)
	fmt.Println("0,1 lower", a, "upper", z)

	a, z = float.SpreadU16(1, 0)
	fmt.Println("1,0 lower", a, "upper", z)

	a, z = float.SpreadU16(1, 1)
	fmt.Println("1,1 lower", a, "upper", z)

	a, z = float.SpreadU16(4, 1)
	fmt.Println("4,1 lower", a, "upper", int(z))

	a, z = float.SpreadU16(5, 1)
	fmt.Println("5,1 lower", a, "upper", int(z))

	a, z = float.SpreadU16(6, 1)
	fmt.Println("6,1 lower", a, "upper", int(z))

	// Output:
	// 0,0 lower 1 upper 1.9999847
	// 0,1 lower 2 upper 3.9999695
	// 1,0 lower 1 upper 3.999939
	// 1,1 lower 2 upper 7.999878
	// 4,1 lower 2 upper 131056
	// 5,1 lower 2 upper 8587837440
	// 6,1 lower 2 upper 9218868437227405312
}
func ExampleSpread16() {
	a, z := float.Spread16(0, 0)
	fmt.Println("0,0 lower", a, "upper", z)

	a, z = float.Spread16(0, 1)
	fmt.Println("0,1 lower", a, "upper", z)

	a, z = float.Spread16(1, 0)
	fmt.Println("1,0 lower", a, "upper", z)

	a, z = float.Spread16(1, 1)
	fmt.Println("1,1 lower", a, "upper", z)

	a, z = float.Spread16(3, 1)
	fmt.Println("3,1 lower", a, "upper", z)

	// Output:
	// 0,0 lower 1 upper 1.9999695
	// 0,1 lower 2 upper 3.999939
	// 1,0 lower 1 upper 3.999878
	// 1,1 lower 2 upper 7.999756
	// 3,1 lower 2 upper 511.9375
}
