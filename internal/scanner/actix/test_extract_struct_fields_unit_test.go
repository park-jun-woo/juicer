//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractStructFields_Unit 테스트
package actix

import "testing"

func TestExtractStructFields_Unit(t *testing.T) {

	src := []byte(`struct Marker;`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	sn := firstStructNode(root)
	if sn == nil {
		t.Fatal("no struct found")
	}
	if fields := extractStructFields(sn, src); fields != nil {
		t.Fatalf("expected nil for unit struct, got %+v", fields)
	}
}
