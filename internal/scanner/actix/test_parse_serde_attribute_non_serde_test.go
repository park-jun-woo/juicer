//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestParseSerdeAttribute_NonSerde 테스트
package actix

import "testing"

func TestParseSerdeAttribute_NonSerde(t *testing.T) {

	src := []byte("#[derive(Deserialize)]\nstruct X { a: i32 }\n")
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	items := collectAttrItems(root)
	if len(items) == 0 {
		t.Fatal("no attribute_item found")
	}
	if sa := parseSerdeAttribute(items[0], src); sa != nil {
		t.Fatalf("expected nil for non-serde attribute, got %+v", sa)
	}
}
