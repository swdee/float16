// Copyright 2019 Montgomery Edwards⁴⁴⁸ and Faye Amacker

package float16_test

import (
	"math"
	"testing"

	"github.com/x448/float16"
)

// prevent compiler optimizing out code by assigning to these
var resultF16 float16.Float16
var resultF32 float32
var resultStr string
var pcn float16.Precision

func BenchmarkFloat32pi(b *testing.B) {
	result := float32(0)
	pi32 := float32(math.Pi)
	pi16 := float16.Fromfloat32(pi32)
	for i := 0; i < b.N; i++ {
		f16 := float16.Frombits(uint16(pi16))
		result = f16.Float32()
	}
	resultF32 = result
}

func BenchmarkFrombits(b *testing.B) {
	result := float16.Float16(0)
	pi32 := float32(math.Pi)
	pi16 := float16.Fromfloat32(pi32)
	for i := 0; i < b.N; i++ {
		result = float16.Frombits(uint16(pi16))
	}
	resultF16 = result
}

func BenchmarkFromFloat32pi(b *testing.B) {
	result := float16.Float16(0)

	pi := float32(math.Pi)
	for i := 0; i < b.N; i++ {
		result = float16.Fromfloat32(pi)
	}
	resultF16 = result
}

func BenchmarkFromFloat32nan(b *testing.B) {
	result := float16.Float16(0)

	nan := float32(math.NaN())
	for i := 0; i < b.N; i++ {
		result = float16.Fromfloat32(nan)
	}
	resultF16 = result
}

func BenchmarkFromFloat32subnorm(b *testing.B) {
	result := float16.Float16(0)

	subnorm := math.Float32frombits(0x007fffff)
	for i := 0; i < b.N; i++ {
		result = float16.Fromfloat32(subnorm)
	}
	resultF16 = result
}

func BenchmarkPrecisionFromFloat32(b *testing.B) {
	var result float16.Precision

	for i := 0; i < b.N; i++ {
		f32 := float32(0.00001) + float32(0.00001)
		result = float16.PrecisionFromfloat32(f32)
	}
	pcn = result
}

func BenchmarkString(b *testing.B) {
	var result string

	pi32 := float32(math.Pi)
	pi16 := float16.Fromfloat32(pi32)
	for i := 0; i < b.N; i++ {
		result = pi16.String()
	}
	resultStr = result
}

func BenchmarkF16toF32LookupConversion(b *testing.B) {
	// Load the buffer outside the loop to avoid reloading it during each iteration.
	float16Buf := loadBuffer()

	b.ResetTimer()

	// Run the benchmark loop.
	for i := 0; i < b.N; i++ {
		float32Buf := make([]float32, len(float16Buf))

		for i, val := range float16Buf {
			float32Buf[i] = float16.FrombitstoF32(val)
		}
	}
}

func BenchmarkF16toF32NormalConversion(b *testing.B) {
	// Load the buffer outside the loop to avoid reloading it during each iteration.
	float16Buf := loadBuffer()

	b.ResetTimer()

	// Run the benchmark loop.
	for i := 0; i < b.N; i++ {
		float32Buf := make([]float32, len(float16Buf))

		for i, val := range float16Buf {
			f16 := float16.Frombits(val)
			float32Buf[i] = f16.Float32()
		}
	}
}

func BenchmarkF16toF32CGOSingleConversion(b *testing.B) {
	// Load the buffer outside the loop to avoid reloading it during each iteration.
	float16Buf := loadBuffer()

	b.ResetTimer()

	// Run the benchmark loop.
	for i := 0; i < b.N; i++ {
		float32Buf := make([]float32, len(float16Buf))

		for i, val := range float16Buf {
			float32Buf[i] = float16.F16tof32single(val)
		}
	}
}

func BenchmarkF16toF32CGOVectorConversion(b *testing.B) {
	// Load the buffer outside the loop to avoid reloading it during each iteration.
	float16Buf := loadBuffer()

	b.ResetTimer()

	// Run the benchmark loop.
	for i := 0; i < b.N; i++ {
		float32Buf := make([]float32, len(float16Buf))

		// Process the buffer in chunks of 8
		for j := 0; j < len(float16Buf); j += 8 {

			// Calculate how many elements are left to avoid out-of-bounds access
			remaining := len(float16Buf) - j

			if remaining >= 8 {
				// Process 8 values at once if there are enough elements left
				float16.F16tof32(
					float16Buf[j],
					float16Buf[j+1],
					float16Buf[j+2],
					float16Buf[j+3],
					float16Buf[j+4],
					float16Buf[j+5],
					float16Buf[j+6],
					float16Buf[j+7],
					&float32Buf[j],
					&float32Buf[j+1],
					&float32Buf[j+2],
					&float32Buf[j+3],
					&float32Buf[j+4],
					&float32Buf[j+5],
					&float32Buf[j+6],
					&float32Buf[j+7],
				)

			} else {
				// process the remaining elements one by one
				//
				// NOTE: our float16Buf loaded does not result in this code
				// being called
				for k := 0; k < remaining; k++ {
					float32Buf[j+k] = float16.F16tof32single(float16Buf[j+k])
				}
			}
		}
	}
}
