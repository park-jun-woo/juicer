//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what Java 타입 → OpenAPI 타입 변환 테스트
package spring

import "testing"

func TestJavaTypeToOpenAPI(t *testing.T) {
	tests := []struct {
		input  string
		typ    string
		format string
		items  string
	}{
		{"String", "string", "", ""},
		{"int", "integer", "int32", ""},
		{"Integer", "integer", "int32", ""},
		{"long", "integer", "int64", ""},
		{"Long", "integer", "int64", ""},
		{"float", "number", "float", ""},
		{"double", "number", "double", ""},
		{"boolean", "boolean", "", ""},
		{"Boolean", "boolean", "", ""},
		{"BigDecimal", "number", "", ""},
		{"LocalDateTime", "string", "date-time", ""},
		{"LocalDate", "string", "date", ""},
		{"UUID", "string", "uuid", ""},
		{"MultipartFile", "string", "binary", ""},
		{"byte[]", "string", "byte", ""},
		{"void", "", "", ""},
		{"List<String>", "array", "", "String"},
		{"Set<Long>", "array", "", "Long"},
		{"Map<String, Object>", "object", "", ""},
		{"UserDto", "object", "", ""},
	}

	for _, tt := range tests {
		oa := javaTypeToOpenAPI(tt.input)
		if oa.Type != tt.typ {
			t.Errorf("%s: type want %s, got %s", tt.input, tt.typ, oa.Type)
		}
		if oa.Format != tt.format {
			t.Errorf("%s: format want %s, got %s", tt.input, tt.format, oa.Format)
		}
		if oa.Items != tt.items {
			t.Errorf("%s: items want %s, got %s", tt.input, tt.items, oa.Items)
		}
	}
}
