package src

import (
	"testing"
)

func TestAverage(t *testing.T) {
	avg := Average(10, 10, 10)
	expected := 11.00

	if avg != expected {
		t.Errorf("The result expected is: %.2f  value returned: %.2f", expected, avg)
	}

	avg = Average(8, 10, 5.6, 3, 2.89)
	expected = 5.90

	if ToFixed(avg, 2) != expected {
		t.Errorf("The result expected is: %.3f  value returned: %.3f", expected, avg)
	}

}
