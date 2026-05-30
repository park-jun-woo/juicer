//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestTryParseBlueprintAssignment_Match 테스트
package flask

import "testing"

func TestTryParseBlueprintAssignment_Match(t *testing.T) {
	assign, b := firstAssignment(t, `api = Blueprint("api", __name__, url_prefix="/api")`+"\n")
	bp := tryParseBlueprintAssignment(assign, b)
	if bp == nil || bp.varName != "api" || bp.name != "api" || bp.urlPrefix != "/api" {
		t.Fatalf("blueprint = %+v", bp)
	}
}
