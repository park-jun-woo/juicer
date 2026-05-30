//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestTryParseBlueprintAssignment_NotCall 테스트
package flask

import "testing"

func TestTryParseBlueprintAssignment_NotCall(t *testing.T) {
	assign, b := firstAssignment(t, "x = 5\n")
	if bp := tryParseBlueprintAssignment(assign, b); bp != nil {
		t.Fatalf("non-call should be nil, got %+v", bp)
	}
}
