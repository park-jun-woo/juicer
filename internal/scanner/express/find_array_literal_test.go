//ff:func feature=scan type=test control=sequence topic=express
//ff:what findArrayLiteral: 매칭+배열 / 이름불일치 / 배열아님 분기
package express

import "testing"

func TestFindArrayLiteral_Found(t *testing.T) {
	fi := mustParse(t, []byte(`const routes = [ {} ]; const other = 1;`))
	arr := findArrayLiteral(fi.Root, fi.Src, "routes")
	if arr == nil || arr.Type() != "array" {
		t.Fatalf("expected array, got %v", arr)
	}
}

func TestFindArrayLiteral_NameMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`const routes = [ {} ];`))
	if arr := findArrayLiteral(fi.Root, fi.Src, "nope"); arr != nil {
		t.Fatalf("expected nil, got %v", arr)
	}
}

func TestFindArrayLiteral_NotArray(t *testing.T) {
	fi := mustParse(t, []byte(`const routes = 42;`))
	if arr := findArrayLiteral(fi.Root, fi.Src, "routes"); arr != nil {
		t.Fatalf("expected nil, got %v", arr)
	}
}
