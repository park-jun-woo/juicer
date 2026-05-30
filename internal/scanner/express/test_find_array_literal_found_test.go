//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindArrayLiteral_Found 테스트
package express

import "testing"

func TestFindArrayLiteral_Found(t *testing.T) {
	fi := mustParse(t, []byte(`const routes = [ {} ]; const other = 1;`))
	arr := findArrayLiteral(fi.Root, fi.Src, "routes")
	if arr == nil || arr.Type() != "array" {
		t.Fatalf("expected array, got %v", arr)
	}
}
