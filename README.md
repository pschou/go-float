# go-float

This package aims to provide a floating point storage method which takes advantage of known details about the number being stored, such as the lower and upper bounds.

By using the SplitAt and JoinAt functions, one can encode and decode a byte slice into a float value with minimal (or no loss) in precision.

Here is an example with 32->24 bits:
```golang
  b := float.Split32BytesAt(3.1415, 4)
  fmt.Println("b =", b[:3])

  flt := float.JoinBytesAt(b[:3], 4)
  fmt.Printf("float %0.6g\n", flt)

  // Output:
  // b = [25 33 202]
  // float 3.1415
```

and again with 64->40 bits:
```golang
  b := float.Split64BytesAt(3.1415926, 4)
  fmt.Println("b =", b[:5])

  flt := float.JoinBytesAt(b[:5], 4)
  fmt.Printf("float %0.8g\n", flt)

  // Output:
  // b = [25 33 251 77 18]
  // float 3.1415926
```
