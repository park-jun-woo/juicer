//ff:func feature=scan type=test control=sequence topic=actix
//ff:what rustComplexTypeToOpenAPI — Vec/Map/Option/기타 변환 분기를 검증
package actix

import "testing"

func TestRustComplexTypeToOpenAPI(t *testing.T) {
	cases := []struct {
		in       string
		wantType string
		wantItem string
	}{
		{"Vec<String>", "array", "String"},
		{"HashMap<String, i32>", "object", ""},
		{"BTreeMap<String, i32>", "object", ""},
		{"Option<i64>", "integer", ""}, // unwraps to inner i64
		{"SomethingElse", "object", ""},
	}
	for _, c := range cases {
		got := rustComplexTypeToOpenAPI(c.in)
		if got.Type != c.wantType {
			t.Errorf("%q: Type = %q, want %q", c.in, got.Type, c.wantType)
		}
		if c.wantItem != "" && got.Items != c.wantItem {
			t.Errorf("%q: Items = %q, want %q", c.in, got.Items, c.wantItem)
		}
	}
}
