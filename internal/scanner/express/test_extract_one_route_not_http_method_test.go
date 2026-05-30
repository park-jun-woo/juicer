//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneRoute_NotHttpMethod 테스트
package express

import "testing"

func TestExtractOneRoute_NotHttpMethod(t *testing.T) {
	fi := mustParse(t, []byte(`r.foo('/x', h);`))
	if ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}); ri != nil {
		t.Fatalf("expected nil for non-http method, got %+v", ri)
	}
}
