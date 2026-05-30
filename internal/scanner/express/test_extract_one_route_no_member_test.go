//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneRoute_NoMember 테스트
package express

import "testing"

func TestExtractOneRoute_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`foo('/x');`))
	if ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}); ri != nil {
		t.Fatalf("expected nil, got %+v", ri)
	}
}
