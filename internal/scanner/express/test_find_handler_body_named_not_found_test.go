//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindHandlerBody_NamedNotFound 테스트
package express

import "testing"

func TestFindHandlerBody_NamedNotFound(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`))
	if body := findHandlerBody(fi, routeInfo{Handler: "missing"}); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
