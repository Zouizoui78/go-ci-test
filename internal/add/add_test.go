package add

import "testing"

func TestAdd(t *testing.T) {
	res := Add(1, 2)
	if res != 3 {
		t.Errorf("expected 3, got %v", res)
	}
}
