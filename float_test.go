package float_test

import (
	"fmt"

	"github.com/pschou/go-float"
)

func ExamplePutScaled32() {
	b := make([]byte, 3)
	float.PutScaled32(b, 3.1415, 4)
	fmt.Println("b =", b)

	flt := float.ReadScaled(b, 4)
	fmt.Printf("float %0.6g\n", flt)

	// Output:
	// b = [25 33 202]
	// float 3.1415
}

func ExamplePutScaled64() {
	b := make([]byte, 5)
	float.PutScaled64(b, 3.1415926, 4)
	fmt.Println("b =", b)

	flt := float.ReadScaled(b, 4)
	fmt.Printf("float %0.8g\n", flt)

	// Output:
	// b = [25 33 251 77 18]
	// float 3.1415926
}

func ExampleReadUScaled32() {
	b := float.UScaled32(3.1415926, 2)

	flt := float.ReadUScaled32(b, 2)
	fmt.Printf("float %0.8g\n", flt)

	// Output:
	// float 3.1415925
}

func ExampleReadUScaled64() {
	b := float.UScaled64(3.1415926, 2)

	flt := float.ReadUScaled64(b, 2)
	fmt.Printf("float %0.8g\n", flt)

	// Output:
	// float 3.1415926
}

func ExampleReadScaled() {
	for _, f := range []float64{106.2345, 23.1234, -17.564, -123.456} {
		b := make([]byte, 4)
		float.PutScaled64(b, f, 8)

		flt := float.ReadScaled(b, 8)
		fmt.Printf("float %0.8g\n", flt)
	}

	// Output:
	// float 106.2345
	// float 23.1234
	// float -17.564
	// float -123.456
}
