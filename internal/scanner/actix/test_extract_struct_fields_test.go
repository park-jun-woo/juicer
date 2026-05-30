//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractStructFields 테스트
package actix

import "testing"

func TestExtractStructFields(t *testing.T) {
	src := []byte(`
struct User {
    id: i64,
    name: String,
}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	sn := firstStructNode(root)
	if sn == nil {
		t.Fatal("no struct found")
	}
	fields := extractStructFields(sn, src)
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d: %+v", len(fields), fields)
	}
	if fields[0].Name != "id" || fields[1].Name != "name" {
		t.Errorf("unexpected fields: %+v", fields)
	}
}
