package float_test

import (
	"fmt"

	"github.com/pschou/go-float"
)

func ExampleSplit32BytesAt() {
	b := float.Split32BytesAt(3.1415, 4)
	fmt.Println("b =", b[:3])

	flt := float.JoinBytesAt(b[:3], 4)
	fmt.Printf("float %0.6g\n", flt)

	// Output:
	// b = [25 33 202]
	// float 3.1415
}

func ExampleSplit64BytesAt() {
	b := float.Split64BytesAt(3.1415926, 4)
	fmt.Println("b =", b[:5])

	flt := float.JoinBytesAt(b[:5], 4)
	fmt.Printf("float %0.8g\n", flt)

	// Output:
	// b = [25 33 251 77 18]
	// float 3.1415926
}

func ExampleSplitBytes() {
	for _, f := range []float64{106.2345, 23.1234, -17.564, -123.456} {
		b := float.Split64BytesAt(f, 8)

		flt := float.JoinBytesAt(b[:4], 8)
		fmt.Printf("float %0.8g\n", flt)
	}

	// Output:
	// float 106.2345
	// float 23.1234
	// float -17.564
	// float -123.456
}

func ExampleOverlap() {
	slice5 := float.Overlap(140.123, 77.456, 8, 7, 5)
	a, b := float.Part(slice5, 8, 7)
	fmt.Printf("5 floatA %g\n", a)
	fmt.Printf("5 floatB %g\n", b)

	slice7 := float.Overlap(140.123, 77.456, 8, 7, 7)
	a, b = float.Part(slice7, 8, 7)
	fmt.Printf("7 floatA %g\n", a)
	fmt.Printf("7 floatB %g\n", b)

	slice9 := float.Overlap(140.123, 77.456, 8, 7, 9)
	a, b = float.Part(slice9, 8, 7)
	fmt.Printf("9 floatA %g\n", a)
	fmt.Printf("9 floatB %g\n", b)

	// Output:
	// 5 floatA 140.12255859375
	// 5 floatB 77.455810546875
	// 7 floatA 140.12299919128418
	// 7 floatB 77.45599937438965
	// 9 floatA 140.12299999594688
	// 9 floatB 77.45599999651313

}
