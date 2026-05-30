//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestSerdeTokenTree_NonSerde 테스트
package actix

import "testing"

func TestSerdeTokenTree_NonSerde(t *testing.T) {

	src := []byte("#[derive(Deserialize)]\nstruct S { a: i32 }\n")
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	items := collectAttrItems(root)
	if len(items) == 0 {
		t.Fatal("no attribute_item")
	}
	if tt := serdeTokenTree(items[0], src); tt != nil {
		t.Fatalf("expected nil for non-serde attr, got %s", tt.Type())
	}
}
