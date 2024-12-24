package float_test

import (
	"fmt"

	"github.com/pschou/go-float"
)

func ExampleToU16() {
	b := float.ToU16(3.1415, 2, 0)
	exp := float.FromU16(b, 2, 0)
	fmt.Println("float", exp)
	// Output:
	// float 3.1414795
}

func ExampleToU16Walk() {
	for _, f := range []float32{100, 1000, 10000, 100000, 1000000, 10000000} {
		b := float.ToU16(f, 3, 12)
		exp := float.FromU16(b, 3, 12)
		fmt.Println("float", exp)
	}
	// Output:
	// float 4096
	// float 4096
	// float 10000
	// float 100000
	// float 1e+06
	// float 1.048512e+06
}

func ExampleTo16() {
	// Fits within 1 to 3.999
	b := float.To16(3.1415, 1, 0)
	exp := float.From16(b, 1, 0)
	fmt.Println("1,0 float", exp)

	// Does not fit within 0.5 to 1.999
	b = float.To16(3.1415, 1, -1)
	exp = float.From16(b, 1, -1)
	fmt.Println("1,-1 float", exp)

	// Fits within 2 to 7.999
	b = float.To16(3.1415, 1, 1)
	exp = float.From16(b, 1, 1)
	fmt.Println("1,1 float", exp)

	// Does not fit within 4 to 15.999
	b = float.To16(3.1415, 1, 2)
	exp = float.From16(b, 1, 2)
	fmt.Println("1,2 float", exp)

	// Output:
	// 1,0 float 3.1414795
	// 1,-1 float 1.999939
	// 1,1 float 3.1414795
	// 1,2 float 4
}

func ExampleLimitsU16() {
	a, z := float.LimitsU16(0, 0)
	fmt.Println("0,0 lower", a, "upper", z)

	a, z = float.LimitsU16(0, 1)
	fmt.Println("0,1 lower", a, "upper", z)

	a, z = float.LimitsU16(1, 0)
	fmt.Println("1,0 lower", a, "upper", z)

	a, z = float.LimitsU16(1, 1)
	fmt.Println("1,1 lower", a, "upper", z)

	a, z = float.LimitsU16(4, 1)
	fmt.Println("4,1 lower", a, "upper", int(z))

	a, z = float.LimitsU16(5, 1)
	fmt.Println("5,1 lower", a, "upper", int(z))

	a, z = float.LimitsU16(6, 1)
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
func ExampleLimits16() {
	a, z := float.Limits16(0, 0)
	fmt.Println("0,0 lower", a, "upper", z)

	a, z = float.Limits16(0, 1)
	fmt.Println("0,1 lower", a, "upper", z)

	a, z = float.Limits16(1, 0)
	fmt.Println("1,0 lower", a, "upper", z)

	a, z = float.Limits16(1, 1)
	fmt.Println("1,1 lower", a, "upper", z)

	a, z = float.Limits16(3, 1)
	fmt.Println("3,1 lower", a, "upper", z)

	// 1 = 0,1  then  15,1
	// 2 = 0..3 then  14,2
	// 3 = 0..7 then  13,3
	// 4 = 0..15 then 12,4
	// 5,11 = 0..2047  then

	// Output:
	// 0,0 lower 1 upper 1.9999695
	// 0,1 lower 2 upper 3.999939
	// 1,0 lower 1 upper 3.999878
	// 1,1 lower 2 upper 7.999756
	// 3,1 lower 2 upper 511.9375
}
