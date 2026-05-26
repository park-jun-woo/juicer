//ff:func feature=scan type=extract control=sequence
//ff:what TestToOpenAPI_WithEndpoints 테스트
package scanner

import "testing"

func TestToOpenAPI_WithEndpoints(t *testing.T) {
	result := &ScanResult{
		Endpoints: []Endpoint{
			{Method: "GET", Path: "/api/users", Handler: "h.ListUsers"},
		},
	}
	data, err := ToOpenAPI(result, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty")
	}
}
