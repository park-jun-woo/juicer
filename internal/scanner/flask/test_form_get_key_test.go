//ff:func feature=scan type=test topic=flask control=sequence
//ff:what formGetKey request.form.get('key') 키 추출 및 비대상 무시 테스트
package flask

import "testing"

func TestFormGetKey(t *testing.T) {
	src := []byte("x = request.form.get('username', '')\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findAllByType(root, "call")[0]
	if got := formGetKey(call, src); got != "username" {
		t.Errorf("got %q", got)
	}
	// not a form.get call
	src2 := []byte("x = other.method('y')\n")
	root2, _ := parsePython(src2)
	call2 := findAllByType(root2, "call")[0]
	if got := formGetKey(call2, src2); got != "" {
		t.Errorf("non-form.get: got %q", got)
	}
}
