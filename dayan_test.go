package yi

import (
	"fmt"
	"strings"
	"testing"
)

// Returns correct DaYan object for valid index within range
func TestReturnsCorrectDaYanForValidIndex(t *testing.T) {
	// Arrr! Let's set sail with a valid index!
	idx := DaYanIndex(10)
	expected := dayanData[(idx-1)%81]

	result := DaYanByIndex(idx)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Handles typical index values correctly
func TestHandlesTypicalIndexValues(t *testing.T) {
	// Ahoy! Testing typical index values!
	idx := DaYanIndex(50)
	expected := dayanData[idx.index()]

	result := DaYanByIndex(idx)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Uses modulo operation to wrap around index values
func TestUsesModuloOperationToWrapAround(t *testing.T) {
	// Shiver me timbers! Testing modulo wrap-around!
	idx := DaYanIndex(82)
	expected := dayanData[idx.index()]

	result := DaYanByIndex(idx)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Handles index value of 0 correctly
func TestHandlesIndexValueOfZero(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if strings.IndexAny(fmt.Sprintf("%v", r), "dayan number out of range") == -1 {
				t.Errorf("Expected panic with message 'dayan number out of range: 0', but got '%v'", r)
			}
		}
	}()
	// Yo-ho-ho! Testing index value of zero!
	idx := DaYanIndex(0)

	DaYanByIndex(idx)

}

// Handles index values equal to DaYanMax correctly
func TestHandlesIndexEqualToDaYanMax(t *testing.T) {
	// Avast! Testing index equal to DaYanMax!
	idx := DaYanIndex(DaYanMax)
	expected := dayanData[idx.index()]

	result := DaYanByIndex(idx)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Handles index values greater than DaYanMax correctly
func TestHandlesIndexGreaterThanDaYanMax(t *testing.T) {
	// Arrr! Testing index greater than DaYanMax!
	idx := DaYanIndex(DaYanMax + 1)
	expected := dayanData[idx.index()]

	result := DaYanByIndex(idx)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Handles maximum possible uint value correctly
func TestHandlesMaxUintValue(t *testing.T) {
	// Blimey! Testing maximum possible uint value!
	idx := ^DaYanIndex(0) // Maximum possible uint value
	expected := dayanData[idx.index()]

	result := DaYanByIndex(idx)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Returns DaYan struct for valid index within range
func TestReturnsDaYanStructForValidIndex(t *testing.T) {
	// Arrr! Let's set sail with a valid index!
	idx := DaYanIndex(10)
	expected := DaYan{Index: 10, Lucky: "Lucky", nvMing: "￥ﾇﾶ", max: true, SkyNine: "SkyNine", Comment: "Comment"}
	dayanData = [DaYanMax]DaYan{9: expected}

	result := DaYanByIndex(idx)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Correctly handles index values that are multiples of 81
func TestHandlesMultiplesOf81(t *testing.T) {
	// Ahoy! Testing multiples of 81!
	idx := DaYanIndex(163)
	expected := DaYan{Index: 1, Lucky: "Lucky", nvMing: "￥ﾇﾶ", max: true, SkyNine: "SkyNine", Comment: "Comment"}
	dayanData = [DaYanMax]DaYan{0: expected}

	result := DaYanByIndex(idx)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Retrieves DaYan struct with accurate field values
func TestRetrievesAccurateFieldValues(t *testing.T) {
	// Yo-ho-ho! Checking accurate field values!
	idx := DaYanIndex(5)
	expected := DaYan{Index: 5, Lucky: "Very Lucky", nvMing: "￥ﾐﾉ", max: false, SkyNine: "Clear Sky", Comment: "No Comment"}
	dayanData = [DaYanMax]DaYan{4: expected}

	result := DaYanByIndex(idx)

	if result.Index != expected.Index || result.Lucky != expected.Lucky || result.nvMing != expected.nvMing || result.max != expected.max || result.SkyNine != expected.SkyNine || result.Comment != expected.Comment {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Manages very large index values by using modulo operation
func TestManagesVeryLargeIndexValues(t *testing.T) {
	// Avast! Testing very large index values!
	idx := DaYanIndex(99999999)
	expected := DaYan{Index: 80, Lucky: "Lucky", nvMing: "￥ﾇﾶ", max: true, SkyNine: "SkyNine", Comment: "Comment"}
	dayanData = [DaYanMax]DaYan{71: expected}

	result := DaYanByIndex(idx)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Returns correct DaYan struct for index values just below 81
func TestReturnsCorrectStructForIndexBelow81(t *testing.T) {
	// Arrr! Testing index just below 81!
	idx := DaYanIndex(80)
	expected := DaYan{Index: 80, Lucky: "Lucky", nvMing: "￥ﾇﾶ", max: true, SkyNine: "SkyNine", Comment: "Comment"}
	dayanData = [DaYanMax]DaYan{79: expected}

	result := DaYanByIndex(idx)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// Returns correct DaYan struct for index values just above 81
func TestReturnsCorrectStructForIndexAbove81(t *testing.T) {
	// Yo-ho-ho! Testing index just above 81!
	idx := DaYanIndex(82)
	expected := DaYan{Index: 1, Lucky: "Lucky", nvMing: "￥ﾇﾶ", max: true, SkyNine: "SkyNine", Comment: "Comment"}
	dayanData = [DaYanMax]DaYan{0: expected}

	result := DaYanByIndex(idx)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
