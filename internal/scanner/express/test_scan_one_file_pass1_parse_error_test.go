//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestScanOneFilePass1_ParseError 테스트
package express

import "testing"

func TestScanOneFilePass1_ParseError(t *testing.T) {
	parsed, allRouters, schemas, schemaSrc := newPass1Maps()
	entries := scanOneFilePass1("/no/such/file.ts", parsed, allRouters, "/no/such", nil, schemas, schemaSrc)
	if entries != nil {
		t.Fatalf("expected nil on parse error, got %+v", entries)
	}
}
