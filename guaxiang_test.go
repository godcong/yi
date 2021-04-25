package yi

import (
	"testing"
)

func TestGetGuaXiang(t *testing.T) {
	if len(getGuaXiangs()) != 64 {
		t.Log("not enough", getGuaXiangs())
	}

}
