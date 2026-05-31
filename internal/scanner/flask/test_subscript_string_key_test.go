//ff:func feature=scan type=test topic=flask control=sequence
//ff:what subscriptStringKey subscript의 문자열 인덱스 추출 테스트
package flask

import "testing"

func TestSubscriptStringKey(t *testing.T) {
	src := []byte("x = request.form['username']\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	sub := findAllByType(root, "subscript")[0]
	if got := subscriptStringKey(sub, src); got != "username" {
		t.Errorf("got %q", got)
	}
	// non-string key
	src2 := []byte("x = arr[0]\n")
	root2, _ := parsePython(src2)
	sub2 := findAllByType(root2, "subscript")[0]
	if got := subscriptStringKey(sub2, src2); got != "" {
		t.Errorf("non-string: got %q", got)
	}
}
