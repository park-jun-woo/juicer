//ff:func feature=scan type=test control=sequence
//ff:what TestInferValueType_CompositeNonArray 테스트
package fiber

import "testing"

func TestInferValueType_CompositeNonArray(t *testing.T) {

	if got := inferFor(t, "Book{}"); got != "unknown" {
		t.Errorf("struct composite -> %q, want unknown", got)
	}
}
