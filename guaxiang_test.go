package yi

import (
	"testing"
)

func TestGetGuaXiang(t *testing.T) {
	if len(GetGuaXiang()) != 64 {
		t.Log("not enough", GetGuaXiang())
	}

}
