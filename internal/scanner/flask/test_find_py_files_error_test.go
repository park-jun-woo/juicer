//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestFindPyFiles_Error 테스트
package flask

import "testing"

func TestFindPyFiles_Error(t *testing.T) {
	if _, err := findPyFiles("/no/such/dir/zzz"); err == nil {
		t.Fatal("expected error for missing root")
	}
}
