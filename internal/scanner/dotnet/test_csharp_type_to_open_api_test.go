//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestCSharpTypeToOpenAPI -- C# 타입 변환 테스트
package dotnet

import "testing"

func TestCSharpTypeToOpenAPI(t *testing.T) {
	tests := []struct {
		input  string
		wType  string
		wFmt   string
		wItems string
	}{
		{"string", "string", "", ""},
		{"int", "integer", "int32", ""},
		{"long", "integer", "int64", ""},
		{"float", "number", "float", ""},
		{"double", "number", "double", ""},
		{"decimal", "number", "", ""},
		{"bool", "boolean", "", ""},
		{"DateTime", "string", "date-time", ""},
		{"DateTimeOffset", "string", "date-time", ""},
		{"DateOnly", "string", "date", ""},
		{"Guid", "string", "uuid", ""},
		{"IFormFile", "string", "binary", ""},
		{"byte[]", "string", "byte", ""},
		{"List<UserDto>", "array", "", "UserDto"},
		{"IEnumerable<string>", "array", "", "string"},
		{"Dictionary<string, int>", "object", "", ""},
		{"UserDto", "object", "", ""},
		{"int?", "integer", "int32", ""},
		{"string?", "string", "", ""},
	}
	for _, tt := range tests {
		oa := csharpTypeToOpenAPI(tt.input)
		if oa.Type != tt.wType {
			t.Errorf("%s: type want %s, got %s", tt.input, tt.wType, oa.Type)
		}
		if oa.Format != tt.wFmt {
			t.Errorf("%s: format want %s, got %s", tt.input, tt.wFmt, oa.Format)
		}
		if oa.Items != tt.wItems {
			t.Errorf("%s: items want %s, got %s", tt.input, tt.wItems, oa.Items)
		}
	}
}
