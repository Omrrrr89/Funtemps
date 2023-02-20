package conv

import (
	"reflect"
	"testing"
	"math"
)
func withinTolerance(a, b, error float64) bool {

    if a == b {
      return true
    }

    difference := math.Abs(a - b)

    if b == 0 {
      return difference < error
    }

 return (difference/math.Abs(b)) < error
}
/*
*

	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
func TestFahrenheitToCelsius(t *testing.T) {

    type test struct {
        input float64
        want  float64
    }

    tests := []test{
        {input: 134, want: 56.67},
    }
    for _, tc := range tests {
        got := FahrenheitToCelsius(tc.input)
        if !withinTolerance(tc.want, got, 1e-3) {
            t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
        }
    }
}

func TestFahrenheitToKevlin(t *testing.T) {
    type test struct {
        input float64
        want  float64
    }

    tests := []test{
        {input: 134, want: 329.82},
    }
    for _, tc := range tests {
        got := FahrenheitToKelvin(tc.input)

        if !withinTolerance(tc.want, got, 1e-2) {
            t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
        }
    }
}
func TestKelvinToCelsius(t *testing.T) {
    type test struct {
        input float64
        want  float64
    }

    tests := []test{
        {input: 329.82, want: 56.67},
    }

    for _, tc := range tests {
        got := KelvinToCelsius(tc.input)

        if !withinTolerance(tc.want, got, 1e-2) {
            t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
        }
    }
}

func TestKelvinToFahrenheit(t *testing.T) {
    type test struct {
        input float64
        want  float64
    }

    tests := []test{
        {input: 329.82, want: 134},
    }

    for _, tc := range tests {
        got := KelvinToFahrenheit(tc.input)

        if !withinTolerance(tc.want, got, 1e-2) {
            t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
        }
    }
}
func TestCelsiusToFahrenheit(t *testing.T) {
    type test struct {
        input float64
        want  float64
    }

    tests := []test{
        {input: 56.67, want: 134},
    }

    for _, tc := range tests {
        got := CelsiusToFahrenheit(tc.input)

        if !withinTolerance(tc.want, got, 1e-2) {
            t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
        }
    }
}

func TestCelsiusToKelvin(t *testing.T) {
    type test struct {
        input float64
        want  float64
    }

    tests := []test{
        {input: 56.67, want: 329.82},
    }

    for _, tc := range tests {
        got := CelsiusToKelvin(tc.input)

        if !withinTolerance(tc.want, got, 1e-2) {
            t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
        }
    }
}
// De andre testfunksjonene implementeres her
// ...
