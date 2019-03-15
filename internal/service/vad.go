package service

import (
	"math"

	"github.com/mjibson/go-dsp/fft"
)

type VAD struct {
	samples      []complex128
	fft          []complex128
	spectrum     []float64
	lastSpectrum []float64
}

func NewVAD(width int) *VAD {
	return &VAD{
		samples:      make([]complex128, 3480),
		spectrum:     make([]float64, 3480/2+1),
		lastSpectrum: make([]float64, 3480/2+1),
	}
}

// Given the samples, return the spectral flux value as compared
// to the previous samples.
func (v *VAD) Flux(samples []int32) float64 {
	for i, s := range samples {
		v.samples[i] = complex(float64(s), 0)
	}

	v.fft = fft.FFT(v.samples)
	copy(v.spectrum, v.lastSpectrum)

	for i, _ := range v.spectrum {
		c := v.fft[i]
		v.spectrum[i] = math.Sqrt(real(c)*real(c) + imag(c)*imag(c))
	}

	var flux float64

	for i, s := range v.spectrum {
		flux += (s - v.lastSpectrum[i])
	}

	return flux
}

func (v *VAD) FFT() []complex128 {
	return v.fft
}
