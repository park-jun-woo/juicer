//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneRoute_NoArgumentsNode 테스트
package express

import "testing"

func TestExtractOneRoute_NoArgumentsNode(t *testing.T) {

	fi := mustParse(t, []byte("r.get`tpl`;"))
	if ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}); ri != nil {
		t.Fatalf("expected nil for no arguments node, got %+v", ri)
	}
}
