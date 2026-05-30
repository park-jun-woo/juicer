//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestScanOneFilePass1_ParseError 테스트
package hono

import "testing"

func TestScanOneFilePass1_ParseError(t *testing.T) {

	if r := scanOneFilePass1("/no/such.ts", "/no"); r != nil {
		t.Fatal("expected nil for parse error")
	}
}
