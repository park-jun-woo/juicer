//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestScanOneFilePass1_ParseError 테스트
package fastify

import "testing"

func TestScanOneFilePass1_ParseError(t *testing.T) {

	if res := scanOneFilePass1("/no/such/file.ts", "/root"); res != nil {
		t.Fatalf("expected nil for unreadable file, got %+v", res)
	}
}
