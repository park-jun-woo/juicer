//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestTryParseBlueprintAssignment_NoUrlPrefix 테스트
package flask

import "testing"

func TestTryParseBlueprintAssignment_NoUrlPrefix(t *testing.T) {
	assign, b := firstAssignment(t, `bp = Blueprint("name", __name__)`+"\n")
	got := tryParseBlueprintAssignment(assign, b)
	if got == nil || got.urlPrefix != "" {
		t.Fatalf("expected empty url_prefix, got %+v", got)
	}
}
