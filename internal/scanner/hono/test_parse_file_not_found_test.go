//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestParseFile_NotFound 테스트
package hono

import "testing"

func TestParseFile_NotFound(t *testing.T) {
	_, err := parseFile("/no/such/file.ts")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}
