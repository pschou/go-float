package float_test

import (
	"fmt"

	"github.com/pschou/go-float"
)

func ExampleUInt16_3() {
	ebits := 3
	b_in := []byte{0, 0x01}
	b_out := []byte{0, 0x01}
	for i := 0; i < 256; i = i + 8 {
		//fmt.Println(ReadInt16([]byte{byte(i >> 8), byte(i)}, 4))
		b_in[0] = byte(i)
		b_in[1] = 0xff
		fi := float.UInt(b_in, ebits)
		float.PutUInt(b_out, uint(fi), ebits)
		fmt.Println(i, b_in, fi, b_out)
	}
	fmt.Println("Maximum", float.UInt([]byte{0xff, 0xff}, ebits))

	// Output:
	// 0 [0 255] 255 [0 255]
	// 8 [8 255] 2303 [8 255]
	// 16 [16 255] 4351 [16 255]
	// 24 [24 255] 6399 [24 255]
	// 32 [32 255] 8447 [32 255]
	// 40 [40 255] 10495 [40 255]
	// 48 [48 255] 12543 [48 255]
	// 56 [56 255] 14591 [56 255]
	// 64 [64 255] 16894 [64 255]
	// 72 [72 255] 20990 [72 255]
	// 80 [80 255] 25086 [80 255]
	// 88 [88 255] 29182 [88 255]
	// 96 [96 255] 33788 [96 255]
	// 104 [104 255] 41980 [104 255]
	// 112 [112 255] 50172 [112 255]
	// 120 [120 255] 58364 [120 255]
	// 128 [128 255] 67576 [128 255]
	// 136 [136 255] 83960 [136 255]
	// 144 [144 255] 100344 [144 255]
	// 152 [152 255] 116728 [152 255]
	// 160 [160 255] 135152 [160 255]
	// 168 [168 255] 167920 [168 255]
	// 176 [176 255] 200688 [176 255]
	// 184 [184 255] 233456 [184 255]
	// 192 [192 255] 270304 [192 255]
	// 200 [200 255] 335840 [200 255]
	// 208 [208 255] 401376 [208 255]
	// 216 [216 255] 466912 [216 255]
	// 224 [224 255] 540608 [224 255]
	// 232 [232 255] 671680 [232 255]
	// 240 [240 255] 802752 [240 255]
	// 248 [248 255] 933824 [248 255]
	// Maximum 1048512
}

func ExampleUInt8_3() {
	ebits := 3
	b_in := []byte{0}
	b_out := []byte{0}
	for i := 0; i < 256; i = i + 8 {
		b_in[0] = byte(i)
		fi := float.UInt(b_in, ebits)
		float.PutUInt(b_out, uint(fi), ebits)
		fmt.Println(i, b_in, fi, b_out)
	}
	fmt.Println("Maximum", float.UInt([]byte{0xff}, ebits))

	// Output:
	// 0 [0] 0 [0]
	// 8 [8] 8 [8]
	// 16 [16] 16 [16]
	// 24 [24] 24 [24]
	// 32 [32] 32 [32]
	// 40 [40] 40 [40]
	// 48 [48] 48 [48]
	// 56 [56] 56 [56]
	// 64 [64] 64 [64]
	// 72 [72] 80 [72]
	// 80 [80] 96 [80]
	// 88 [88] 112 [88]
	// 96 [96] 128 [96]
	// 104 [104] 160 [104]
	// 112 [112] 192 [112]
	// 120 [120] 224 [120]
	// 128 [128] 256 [128]
	// 136 [136] 320 [136]
	// 144 [144] 384 [144]
	// 152 [152] 448 [152]
	// 160 [160] 512 [160]
	// 168 [168] 640 [168]
	// 176 [176] 768 [176]
	// 184 [184] 896 [184]
	// 192 [192] 1024 [192]
	// 200 [200] 1280 [200]
	// 208 [208] 1536 [208]
	// 216 [216] 1792 [216]
	// 224 [224] 2048 [224]
	// 232 [232] 2560 [232]
	// 240 [240] 3072 [240]
	// 248 [248] 3584 [248]
	// Maximum 4032
}

func ExampleUInt16_4() {
	ebits := 4
	b_in := []byte{0, 0x01}
	b_out := []byte{0, 0x01}
	for i := 0; i < 256; i = i + 8 {
		//fmt.Println(ReadInt16([]byte{byte(i >> 8), byte(i)}, 4))
		b_in[0] = byte(i)
		b_in[1] = 0xff
		fi := float.UInt(b_in, ebits)
		float.PutUInt(b_out, uint(fi), ebits)
		fmt.Println(i, b_in, fi, b_out)
	}
	fmt.Println("Maximum", float.UInt([]byte{0xff, 0xff}, ebits))

	// Output:
	// 0 [0 255] 255 [0 255]
	// 8 [8 255] 2303 [8 255]
	// 16 [16 255] 4351 [16 255]
	// 24 [24 255] 6399 [24 255]
	// 32 [32 255] 8702 [32 255]
	// 40 [40 255] 12798 [40 255]
	// 48 [48 255] 17404 [48 255]
	// 56 [56 255] 25596 [56 255]
	// 64 [64 255] 34808 [64 255]
	// 72 [72 255] 51192 [72 255]
	// 80 [80 255] 69616 [80 255]
	// 88 [88 255] 102384 [88 255]
	// 96 [96 255] 139232 [96 255]
	// 104 [104 255] 204768 [104 255]
	// 112 [112 255] 278464 [112 255]
	// 120 [120 255] 409536 [120 255]
	// 128 [128 255] 556928 [128 255]
	// 136 [136 255] 819072 [136 255]
	// 144 [144 255] 1113856 [144 255]
	// 152 [152 255] 1638144 [152 255]
	// 160 [160 255] 2227712 [160 255]
	// 168 [168 255] 3276288 [168 255]
	// 176 [176 255] 4455424 [176 255]
	// 184 [184 255] 6552576 [184 255]
	// 192 [192 255] 8910848 [192 255]
	// 200 [200 255] 13105152 [200 255]
	// 208 [208 255] 17821696 [208 255]
	// 216 [216 255] 26210304 [216 255]
	// 224 [224 255] 35643392 [224 255]
	// 232 [232 255] 52420608 [232 255]
	// 240 [240 255] 71286784 [240 255]
	// 248 [248 255] 104841216 [248 255]
	// Maximum 134201344
}

func ExampleInt16_3() {
	ebits := 3
	b_in := []byte{0, 0x01}
	b_out := []byte{0, 0x01}
	for i := 0; i < 256; i = i + 8 {
		//fmt.Println(ReadInt16([]byte{byte(i >> 8), byte(i)}, 4))
		b_in[0] = byte(i)
		b_in[1] = 0xff
		fi := float.Int(b_in, ebits)
		float.PutInt(b_out, fi, ebits)
		fmt.Println(i, b_in, fi, b_out)
	}
	fmt.Println("Maximum", float.Int([]byte{0x7f, 0xff}, ebits))

	// Output:
	// 0 [0 255] 255 [0 255]
	// 8 [8 255] 2303 [8 255]
	// 16 [16 255] 4351 [16 255]
	// 24 [24 255] 6399 [24 255]
	// 32 [32 255] 8702 [32 255]
	// 40 [40 255] 12798 [40 255]
	// 48 [48 255] 17404 [48 255]
	// 56 [56 255] 25596 [56 255]
	// 64 [64 255] 34808 [64 255]
	// 72 [72 255] 51192 [72 255]
	// 80 [80 255] 69616 [80 255]
	// 88 [88 255] 102384 [88 255]
	// 96 [96 255] 139232 [96 255]
	// 104 [104 255] 204768 [104 255]
	// 112 [112 255] 278464 [112 255]
	// 120 [120 255] 409536 [120 255]
	// 128 [128 255] -255 [128 255]
	// 136 [136 255] -2303 [136 255]
	// 144 [144 255] -4351 [144 255]
	// 152 [152 255] -6399 [152 255]
	// 160 [160 255] -8702 [160 255]
	// 168 [168 255] -12798 [168 255]
	// 176 [176 255] -17404 [176 255]
	// 184 [184 255] -25596 [184 255]
	// 192 [192 255] -34808 [192 255]
	// 200 [200 255] -51192 [200 255]
	// 208 [208 255] -69616 [208 255]
	// 216 [216 255] -102384 [216 255]
	// 224 [224 255] -139232 [224 255]
	// 232 [232 255] -204768 [232 255]
	// 240 [240 255] -278464 [240 255]
	// 248 [248 255] -409536 [248 255]
	// Maximum 524224
}
