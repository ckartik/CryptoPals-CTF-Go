package hammingdist

import "testing"

func TestCalculateDistance(t *testing.T) {
	result := CalculateDistance("this is a test", "wokka wokka!!!")

	if result != 37 {
		t.Errorf("Error: Expcted 37 got %v.", result)
	}
}
