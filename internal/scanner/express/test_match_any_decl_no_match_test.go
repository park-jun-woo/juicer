//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchAnyDecl_NoMatch 테스트
package express

import "testing"

func TestMatchAnyDecl_NoMatch(t *testing.T) {
	fi := mustParse(t, []byte(`function handler() {}`))
	if body := matchAnyDecl(topChild(t, fi, "function_declaration"), fi.Src, "other"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
