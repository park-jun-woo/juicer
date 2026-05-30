//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestParseFile_Missing 테스트
package supafunc

import "testing"

func TestParseFile_Missing(t *testing.T) {
	if _, err := parseFile("/no/such.ts"); err == nil {
		t.Fatal("expected error")
	}
}
