//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestGoTypeToOpenAPI_String 타입 변환 반복 테스트
package scanner

import "testing"

func TestGoTypeToOpenAPI_String(t *testing.T) {
	cases := []struct{ in, out string }{
		{"string", "string"},
		{"*string", "string"},
		{"int", "integer"},
		{"int64", "integer"},
		{"uint32", "integer"},
		{"float32", "number"},
		{"float64", "number"},
		{"bool", "boolean"},
		{"boolean", "boolean"},
		{"integer", "integer"},
		{"number", "number"},
		{"object", "object"},
		{"array", "array"},
		{"any", "object"},
		{"interface{}", "object"},
		{"[]string", "array"},
		{"time.Time", "string"},
		{"CustomType", "object"},
	}
	for _, c := range cases {
		if got := goTypeToOpenAPI(c.in); got != c.out {
			t.Fatalf("goTypeToOpenAPI(%q)=%q, want %q", c.in, got, c.out)
		}
	}
}

