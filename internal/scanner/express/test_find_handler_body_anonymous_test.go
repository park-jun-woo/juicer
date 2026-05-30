//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindHandlerBody_Anonymous 테스트
package express

import "testing"

func TestFindHandlerBody_Anonymous(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`))
	if body := findHandlerBody(fi, routeInfo{Handler: "(anonymous)"}); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
