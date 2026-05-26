//ff:func feature=scan type=test control=iteration dimension=1
//ff:what buildParamSchema 테스트
package scanner

import "testing"

func TestBuildParamSchema(t *testing.T) {
	cases := []struct {
		typ        string
		wantType   string
		wantFormat string
		hasFormat  bool
	}{
		{"integer", "integer", "", false},
		{"string:uuid", "string", "uuid", true},
		{"string:date-time", "string", "date-time", true},
		{"", "string", "", false},
		{"number", "number", "", false},
	}
	for _, c := range cases {
		got := buildParamSchema(c.typ)
		if got["type"] != c.wantType {
			t.Errorf("buildParamSchema(%q)[type] = %q, want %q", c.typ, got["type"], c.wantType)
		}
		gotFmt, hasFmt := got["format"]
		if hasFmt != c.hasFormat {
			t.Errorf("buildParamSchema(%q) format present=%v, want %v", c.typ, hasFmt, c.hasFormat)
		}
		if c.hasFormat && gotFmt != c.wantFormat {
			t.Errorf("buildParamSchema(%q)[format] = %q, want %q", c.typ, gotFmt, c.wantFormat)
		}
	}
}
