package text

import "testing"

func TestIsPhone(t *testing.T) {
	if IsPhone("13358454751a") {
		t.Failed()
	}
	if !IsPhone("13025928479") {
		t.Failed()
	}
}