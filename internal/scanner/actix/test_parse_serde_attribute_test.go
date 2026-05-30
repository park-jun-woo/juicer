//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestParseSerdeAttribute 테스트
package actix

import "testing"

func TestParseSerdeAttribute(t *testing.T) {
	root, err := parseRust([]byte(serdeAttrsSource))
	if err != nil {
		t.Fatal(err)
	}
	var renamed, def, skip bool
	for _, item := range collectAttrItems(root) {
		sa := parseSerdeAttribute(item, []byte(serdeAttrsSource))
		if sa == nil {
			continue
		}
		if sa.rename == "userName" {
			renamed = true
		}
		if sa.hasDefault {
			def = true
		}
		if sa.skip {
			skip = true
		}
	}
	if !renamed {
		t.Error("expected a serde rename = userName to be parsed")
	}
	if !def {
		t.Error("expected a serde default to be parsed")
	}
	if !skip {
		t.Error("expected a serde skip to be parsed")
	}
}
