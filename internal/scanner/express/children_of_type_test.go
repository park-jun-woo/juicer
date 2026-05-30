//ff:func feature=scan type=test control=sequence topic=express
//ff:what childrenOfType — 직접 자식 중 지정 타입 수집을 검증
package express

import "testing"

func TestChildrenOfType(t *testing.T) {
	fi := mustParse(t, []byte("const a = [1, 2, 3];\n"))
	arr := firstParamOfType(fi.Root, "array")
	if arr == nil {
		t.Fatal("no array node")
	}
	nums := childrenOfType(arr, "number")
	if len(nums) != 3 {
		t.Fatalf("expected 3 numbers, got %d", len(nums))
	}
	if got := childrenOfType(arr, "string"); len(got) != 0 {
		t.Fatalf("expected 0 strings, got %d", len(got))
	}
}
