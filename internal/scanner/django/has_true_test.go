//ff:func feature=scan type=test control=sequence topic=django
//ff:what hasTrue — keyword 인자에 True 값 존재 여부를 검증
package django

import "testing"

func TestHasTrue(t *testing.T) {
	srcTrue := []byte("x = action(detail=True)\n")
	root, err := parsePython(srcTrue)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if len(kw) == 0 {
		t.Fatal("no keyword_argument")
	}
	if !hasTrue(kw[0], srcTrue) {
		t.Error("expected hasTrue true for detail=True")
	}

	srcFalse := []byte("x = action(detail=False)\n")
	root2, err := parsePython(srcFalse)
	if err != nil {
		t.Fatal(err)
	}
	kw2 := keywordArgs(root2)
	if hasTrue(kw2[0], srcFalse) {
		t.Error("expected hasTrue false for detail=False")
	}
}
