//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestParseFile_ReadError 테스트
package fastify

import "testing"

func TestParseFile_ReadError(t *testing.T) {
	if _, err := parseFile("/no/such/file.ts"); err == nil {
		t.Fatal("expected read error")
	}
}
