//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestPyTypeToOpenAPI_List 테스트
package fastapi

import "testing"

func TestPyTypeToOpenAPI_List(t *testing.T) {
	tests := []struct {
		py        string
		wantType  string
		wantItems string
	}{
		{"list[str]", "array", "str"},
		{"List[int]", "array", "int"},
	}
	for _, tc := range tests {
		oa := pyTypeToOpenAPI(tc.py)
		if oa.Type != tc.wantType {
			t.Errorf("pyTypeToOpenAPI(%q).Type = %q, want %q", tc.py, oa.Type, tc.wantType)
		}
		if oa.Items != tc.wantItems {
			t.Errorf("pyTypeToOpenAPI(%q).Items = %q, want %q", tc.py, oa.Items, tc.wantItems)
		}
	}
}
