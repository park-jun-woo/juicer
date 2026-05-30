//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindInitValue_None 테스트
package express

import "testing"

func TestFindInitValue_None(t *testing.T) {
	fi := mustParse(t, []byte(`const r = 42;`))
	if v := findInitValue(firstDeclarator(t, fi)); v != nil {
		t.Fatalf("expected nil, got %v", v.Type())
	}
}
