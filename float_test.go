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
