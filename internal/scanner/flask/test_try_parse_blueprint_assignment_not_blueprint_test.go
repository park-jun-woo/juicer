//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestTryParseBlueprintAssignment_NotBlueprint 테스트
package flask

import "testing"

func TestTryParseBlueprintAssignment_NotBlueprint(t *testing.T) {
	assign, b := firstAssignment(t, `app = Flask(__name__)`+"\n")
	if bp := tryParseBlueprintAssignment(assign, b); bp != nil {
		t.Fatalf("Flask() should be nil, got %+v", bp)
	}
}
