//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneRoute_NoArgsBuildNil 테스트
package express

import "testing"

func TestExtractOneRoute_NoArgsBuildNil(t *testing.T) {

	fi := mustParse(t, []byte(`r.get();`))
	if ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}); ri != nil {
		t.Fatalf("expected nil for no args, got %+v", ri)
	}
}
