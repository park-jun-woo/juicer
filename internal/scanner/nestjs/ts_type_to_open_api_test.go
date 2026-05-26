//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what tsTypeToOpenAPI 테스트
package nestjs

import "testing"

func TestTsTypeToOpenAPI_Cases(t *testing.T) {
	cases := []struct {
		in         string
		wantType   string
		wantFormat string
		wantItems  string
	}{
		{"string", "string", "", ""},
		{"number", "number", "", ""},
		{"boolean", "boolean", "", ""},
		{"Date", "string", "date-time", ""},
		{"Uuid", "string", "uuid", ""},
		{"ObjectId", "string", "", ""},
		{"any", "object", "", ""},
		{"void", "", "", ""},
		{"", "", "", ""},
		{"Array<string>", "array", "", "string"},
		{"number[]", "array", "", "number"},
		{"Record<string, number>", "object", "", "number"},
		{"Record<string>", "object", "", ""},
		{"MyDto", "object", "", ""},
		{"Promise<string>", "string", "", ""},
	}
	for _, c := range cases {
		got := tsTypeToOpenAPI(c.in)
		if got.Type != c.wantType {
			t.Errorf("tsTypeToOpenAPI(%q).Type = %q, want %q", c.in, got.Type, c.wantType)
		}
		if got.Format != c.wantFormat {
			t.Errorf("tsTypeToOpenAPI(%q).Format = %q, want %q", c.in, got.Format, c.wantFormat)
		}
		if got.Items != c.wantItems {
			t.Errorf("tsTypeToOpenAPI(%q).Items = %q, want %q", c.in, got.Items, c.wantItems)
		}
	}
}
