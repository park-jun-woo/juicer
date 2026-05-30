//ff:func feature=scan type=test control=sequence topic=actix
//ff:what serdeTokenTree — serde 어트리뷰트 token_tree 탐색 분기를 검증
package actix

import "testing"

func TestSerdeTokenTree_Serde(t *testing.T) {
	src := []byte("struct S {\n #[serde(rename = \"x\")]\n a: String,\n}\n")
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	items := collectAttrItems(root)
	var got bool
	for _, it := range items {
		if tt := serdeTokenTree(it, src); tt != nil {
			got = true
			if tt.Type() != "token_tree" {
				t.Errorf("expected token_tree, got %s", tt.Type())
			}
		}
	}
	if !got {
		t.Error("expected serde token_tree to be found")
	}
}

func TestSerdeTokenTree_NonSerde(t *testing.T) {
	// #[derive(Deserialize)] -> identifier name is "derive", not "serde" -> nil.
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
