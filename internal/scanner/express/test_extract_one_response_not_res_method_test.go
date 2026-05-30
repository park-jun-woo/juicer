//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneResponse_NotResMethod 테스트
package express

import "testing"

func TestExtractOneResponse_NotResMethod(t *testing.T) {
	fi := mustParse(t, []byte(`foo.bar(1);`))
	r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
	if r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
