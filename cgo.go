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
