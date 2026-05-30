//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestParseFile_ReadError 테스트
package flask

import "testing"

func TestParseFile_ReadError(t *testing.T) {
	if _, err := parseFile("/root", "/no/such/file.py"); err == nil {
		t.Fatal("expected read error")
	}
}
