package gdsp

import (
	"math/cmplx"
	"testing"
)

func BenchmarkFFT(b *testing.B) {
	v := []complex128{1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0}

	for i := 0; i < b.N; i++ {
		FFT(v)
	}
}

func BenchmarkInverseDFT(b *testing.B) {
	v := []complex128{1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0}

	for i := 0; i < b.N; i++ {
		DFT(v, false)
	}
}

func BenchmarkInverseFFT(b *testing.B) {
	v := []complex128{1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0}

	for i := 0; i < b.N; i++ {
		IFFT(v)
	}
}

func BenchmarkDFT(b *testing.B) {
	v := []complex128{1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0}

	for i := 0; i < b.N; i++ {
		DFT(v, true)
	}
}

func TestDFT(t *testing.T) {
	v := []complex128{1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0}
	dft := DFT(v, true)
	fft := FFT(v)

	for i := 0; i < len(dft); i++ {
		if cmplx.Abs(dft[i]-fft[i]) > 0.000001 {
			t.Errorf("DFT value %f at index %d is not near FFT value of %f.", dft[i], i, fft[i])
		}
	}
}

func TestInverseDFT(t *testing.T) {
	v := []complex128{1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0, 1.0, 2.0, 3.0, 4.0, 4.0, 3.0, 2.0, 1.0}
	dft := DFT(v, false)
	fft := IFFT(v)

	for i := 0; i < len(dft); i++ {
		if cmplx.Abs(dft[i]-fft[i]) > 0.000001 {
			t.Fail()
		}
	}
}
