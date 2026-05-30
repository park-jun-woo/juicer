//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindNamedFunctionBody_NotFound 테스트
package express

import "testing"

func TestFindNamedFunctionBody_NotFound(t *testing.T) {
	fi := mustParse(t, []byte(`function getUsers() {}`))
	if body := findNamedFunctionBody(fi, "missing"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
