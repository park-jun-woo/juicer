//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestTryParseBlueprintAssignment_NoLeftIdentifier 테스트
package flask

import "testing"

func TestTryParseBlueprintAssignment_NoLeftIdentifier(t *testing.T) {

	assign, b := firstAssignment(t, `obj.attr = Blueprint("x", __name__)`+"\n")
	if bp := tryParseBlueprintAssignment(assign, b); bp != nil {
		t.Fatalf("attribute LHS should be nil, got %+v", bp)
	}
}
