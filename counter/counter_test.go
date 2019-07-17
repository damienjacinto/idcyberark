package counter

import (
	"testing"
)

func TestCounter(t *testing.T) {
	counterSafe := New(3)
	if have, want := counterSafe.max, 3; have != want {
		t.Errorf("Init val of counterSafe is wrong. Have: %d, want: %d.", have, want)
	}

	id1 := counterSafe.Inc("test")

	if have, want := id1, 1; have != want {
		t.Errorf("Inc of counter failed. Have: %d, want: %d.", have, want)
	}

	id2 := counterSafe.Inc("test")

	if have, want := id2, 2; have != want {
		t.Errorf("Inc of counter failed. Have: %d, want: %d.", have, want)
	}
}
