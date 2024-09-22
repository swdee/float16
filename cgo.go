package float16

/*
#cgo CFLAGS: -march=native -mtune=native -Ofast -flto
#cgo LDFLAGS: -march=native -mtune=native -Ofast

#include <stdint.h>

// Optimized out into a single vector instruction on x86 compatibles and aarch64
void float16_to_float32(
    uint16_t a, uint16_t b, uint16_t c, uint16_t d,
    uint16_t e, uint16_t f, uint16_t g, uint16_t h,
    float* res_a, float* res_b, float* res_c, float* res_d,
    float* res_e, float* res_f, float* res_g, float* res_h)
{
    _Float16 tmp_a = *(_Float16*)&a;
    _Float16 tmp_b = *(_Float16*)&b;
    _Float16 tmp_c = *(_Float16*)&c;
    _Float16 tmp_d = *(_Float16*)&d;
    _Float16 tmp_e = *(_Float16*)&e;
    _Float16 tmp_f = *(_Float16*)&f;
    _Float16 tmp_g = *(_Float16*)&g;
    _Float16 tmp_h = *(_Float16*)&h;
    *res_a = (float)tmp_a;
    *res_b = (float)tmp_b;
    *res_c = (float)tmp_c;
    *res_d = (float)tmp_d;
    *res_e = (float)tmp_e;
    *res_f = (float)tmp_f;
    *res_g = (float)tmp_g;
    *res_h = (float)tmp_h;
}

void float16_to_float32_single(uint16_t a, float* res) {
    _Float16 tmp = *(_Float16*)&a;
    *res = (float)tmp;
}

void float16_to_float32_buffer(const uint16_t* input, float* output, size_t count) {
    for (size_t i = 0; i < count; i++) {
        _Float16 tmp = *(_Float16*)&input[i];
        output[i] = (float)tmp;
    }
}

void float16_to_float32_vector_buffer(const uint16_t* input, float* output, size_t count) {
    // Process in chunks of 8
    for (size_t i = 0; i < count; i += 8) {
        if (i + 8 <= count) {
            // Process 8 elements in one go
            _Float16 tmp_a = *(_Float16*)&input[i];
            _Float16 tmp_b = *(_Float16*)&input[i + 1];
            _Float16 tmp_c = *(_Float16*)&input[i + 2];
            _Float16 tmp_d = *(_Float16*)&input[i + 3];
            _Float16 tmp_e = *(_Float16*)&input[i + 4];
            _Float16 tmp_f = *(_Float16*)&input[i + 5];
            _Float16 tmp_g = *(_Float16*)&input[i + 6];
            _Float16 tmp_h = *(_Float16*)&input[i + 7];
            output[i] = (float)tmp_a;
            output[i + 1] = (float)tmp_b;
            output[i + 2] = (float)tmp_c;
            output[i + 3] = (float)tmp_d;
            output[i + 4] = (float)tmp_e;
            output[i + 5] = (float)tmp_f;
            output[i + 6] = (float)tmp_g;
            output[i + 7] = (float)tmp_h;
        } else {
            // Process remaining elements one by one
            for (size_t j = i; j < count; j++) {
                _Float16 tmp = *(_Float16*)&input[j];
                output[j] = (float)tmp;
            }
        }
    }
}


*/
import "C"
import "unsafe"

func F16tof32(
	a, b, c, d, e, f, g, h uint16,
	resA, resB, resC, resD, resE, resF, resG, resH *float32) {

	C.float16_to_float32(
		C.uint16_t(a), C.uint16_t(b), C.uint16_t(c), C.uint16_t(d),
		C.uint16_t(e), C.uint16_t(f), C.uint16_t(g), C.uint16_t(h),
		(*C.float)(resA), (*C.float)(resB), (*C.float)(resC), (*C.float)(resD),
		(*C.float)(resE), (*C.float)(resF), (*C.float)(resG), (*C.float)(resH),
	)
}

// F16tof32 converts a single uint16 to float32
func F16tof32single(a uint16) float32 {
	var res float32
	C.float16_to_float32_single(C.uint16_t(a), (*C.float)(unsafe.Pointer(&res)))
	return res
}

// F16toF32Buffer converts an entire float16 buffer to float32 buffer
func F16toF32BufferSingle(float16Buf []uint16, float32Buf []float32) {
	C.float16_to_float32_buffer(
		(*C.uint16_t)(unsafe.Pointer(&float16Buf[0])), // Pointer to the input buffer
		(*C.float)(unsafe.Pointer(&float32Buf[0])),    // Pointer to the output buffer
		C.size_t(len(float16Buf)),                     // Number of elements to convert
	)
}

func F16toF32BufferVector(float16Buf []uint16, float32Buf []float32) {
	C.float16_to_float32_vector_buffer(
		(*C.uint16_t)(unsafe.Pointer(&float16Buf[0])), // Pointer to the input buffer
		(*C.float)(unsafe.Pointer(&float32Buf[0])),    // Pointer to the output buffer
		C.size_t(len(float16Buf)),                     // Number of elements to convert
	)
}
