//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestSerdeTokenTree_Serde 테스트
package actix

import "testing"

func TestSerdeTokenTree_Serde(t *testing.T) {
	src := []byte("struct S {\n #[serde(rename = \"x\")]\n a: String,\n}\n")
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	items := collectAttrItems(root)
	tt := findSerdeTokenTree(items, src)
	if tt == nil {
		t.Error("expected serde token_tree to be found")
	} else if tt.Type() != "token_tree" {
		t.Errorf("expected token_tree, got %s", tt.Type())
	}
}
