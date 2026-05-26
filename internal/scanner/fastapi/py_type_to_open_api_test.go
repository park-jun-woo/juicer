//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestPyTypeToOpenAPI_Builtin 테스트
package fastapi

import "testing"

func TestPyTypeToOpenAPI_Builtin(t *testing.T) {
	tests := []struct {
		py       string
		wantType string
		wantFmt  string
	}{
		{"str", "string", ""},
		{"int", "integer", ""},
		{"float", "number", ""},
		{"bool", "boolean", ""},
		{"dict", "object", ""},
		{"datetime", "string", "date-time"},
		{"date", "string", "date"},
		{"EmailStr", "string", "email"},
		{"uuid.UUID", "string", "uuid"},
		{"UUID", "string", "uuid"},
		{"ObjectId", "string", ""},
		{"PydanticObjectId", "string", ""},
		{"bson.ObjectId", "string", ""},
		{"odmantic.bson.ObjectId", "string", ""},
		{"Any", "object", ""},
	}
	for _, tc := range tests {
		oa := pyTypeToOpenAPI(tc.py)
		if oa.Type != tc.wantType {
			t.Errorf("pyTypeToOpenAPI(%q).Type = %q, want %q", tc.py, oa.Type, tc.wantType)
		}
		if oa.Format != tc.wantFmt {
			t.Errorf("pyTypeToOpenAPI(%q).Format = %q, want %q", tc.py, oa.Format, tc.wantFmt)
		}
	}
}
