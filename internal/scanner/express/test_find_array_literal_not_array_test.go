//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindArrayLiteral_NotArray 테스트
package express

import "testing"

func TestFindArrayLiteral_NotArray(t *testing.T) {
	fi := mustParse(t, []byte(`const routes = 42;`))
	if arr := findArrayLiteral(fi.Root, fi.Src, "routes"); arr != nil {
		t.Fatalf("expected nil, got %v", arr)
	}
}
