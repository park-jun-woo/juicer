//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneRoute_NoIdentifierObject 테스트
package express

import "testing"

func TestExtractOneRoute_NoIdentifierObject(t *testing.T) {

	fi := mustParse(t, []byte(`a.b.get('/x', h);`))
	if ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"a": true, "b": true}); ri != nil {
		t.Fatalf("expected nil, got %+v", ri)
	}
}
