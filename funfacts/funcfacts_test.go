package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string // her må du skrive riktig type for input
		want []string  // her må du skrive riktig type for returverdien
	}

	// Her må du legge inn korrekte testverdier
	tests := []test{
		{input: "Sun", want: []string{ "The sun makes up 99.86% of the mass in the solar system.",
      "The sun is actually white, not yellow.",
      "It takes 8 minutes and 20 seconds for light to travel from the sun to Earth.",}},
		{input: "Luna", want: []string{"The moon is not a perfect sphere, but is slightly flattened at the poles and bulges at the equator.",
		"The moon's gravity is about one-sixth that of Earth's gravity.",
		"The moon is moving away from the Earth at a rate of about 1.5 inches per year.",}},
		{input: "Terra", want: []string{"Earth is the only planet known to support life.",
		"70% of the Earth's surface is covered in water.",
		"The tallest mountain on Earth, Mount Everest, is 29,029 feet tall.",}},

	}


	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
