//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what rustTypeToOpenAPI — 모든 스칼라/포맷/복합 타입 변환 분기를 검증
package actix

import "testing"

func TestRustTypeToOpenAPI_AllBranches(t *testing.T) {
	cases := []struct {
		in         string
		wantType   string
		wantFormat string
	}{
		{"  String  ", "string", ""}, // trimmed
		{"str", "string", ""},
		{"&str", "string", ""},
		{"i8", "integer", "int32"},
		{"i16", "integer", "int32"},
		{"i32", "integer", "int32"},
		{"i64", "integer", "int64"},
		{"u8", "integer", "int32"},
		{"u16", "integer", "int32"},
		{"u32", "integer", "int32"},
		{"u64", "integer", "int64"},
		{"f32", "number", "float"},
		{"f64", "number", "double"},
		{"bool", "boolean", ""},
		{"Uuid", "string", "uuid"},
		{"NaiveDateTime", "string", "date-time"},
		{"DateTime", "string", "date-time"},
		{"DateTime<Utc>", "string", "date-time"},
		{"NaiveDate", "string", "date"},
		{"Vec<i32>", "array", ""},
		{"MyStruct", "object", ""},
	}
	for _, c := range cases {
		got := rustTypeToOpenAPI(c.in)
		if got.Type != c.wantType {
			t.Errorf("%q: Type = %q, want %q", c.in, got.Type, c.wantType)
		}
		if got.Format != c.wantFormat {
			t.Errorf("%q: Format = %q, want %q", c.in, got.Format, c.wantFormat)
		}
	}
}
