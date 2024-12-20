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
	b := float.Squeeze16(3.1415, 1, 1)
	exp := float.Expand16(b, 1, 1)
	fmt.Println("1,1 float", exp)

	// Does not fit within 0.5 to 1.999
	b = float.Squeeze16(3.1415, 1, 0)
	exp = float.Expand16(b, 1, 0)
	fmt.Println("1,0 float", exp)

	// Fits within 2 to 7.999
	b = float.Squeeze16(3.1415, 1, 2)
	exp = float.Expand16(b, 1, 2)
	fmt.Println("1,2 float", exp)

	// Does not fit within 4 to 15.999
	b = float.Squeeze16(3.1415, 1, 3)
	exp = float.Expand16(b, 1, 3)
	fmt.Println("1,3 float", exp)

	// Output:
	// 1,1 float 3.1414795
	// 1,0 float 1.999939
	// 1,2 float 3.1414795
	// 1,3 float 4
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

	a, z = float.SpreadU16(3, 1)
	fmt.Println("0,2 lower", a, "upper", z)

	// Output:
	// 0,0 lower 0.5 upper 0.9999924
	// 0,1 lower 1 upper 1.9999847
	// 1,0 lower 0.5 upper 1.9999695
	// 1,1 lower 1 upper 3.999939
	// 0,2 lower 1 upper 255.98438
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
	fmt.Println("0,2 lower", a, "upper", z)

	// Output:
	// 0,0 lower 0.5 upper 0.99998474
	// 0,1 lower 1 upper 1.9999695
	// 1,0 lower 0.5 upper 1.999939
	// 1,1 lower 1 upper 3.999878
	// 0,2 lower 1 upper 255.96875
}
