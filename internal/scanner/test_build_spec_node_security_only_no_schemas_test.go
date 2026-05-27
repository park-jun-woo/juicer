//ff:func feature=scan type=test control=sequence
//ff:what TestBuildSpecNode_SecurityOnlyNoSchemas 테스트
package scanner

import "testing"

func TestBuildSpecNode_SecurityOnlyNoSchemas(t *testing.T) {
	// auth endpoint but no body/response schemas -> components should still exist with securitySchemes
	result := &ScanResult{
		Endpoints: []Endpoint{
			{
				Method:    "GET",
				Path:      "/api/me",
				AuthLevel: "auth_required",
			},
		},
	}
	node := buildSpecNode(result)
	comp := findMappingValue(node, "components")
	if comp == nil {
		t.Fatal("expected components even without schemas")
	}
	schemas := findMappingValue(comp, "schemas")
	if schemas != nil {
		t.Fatal("expected no schemas node")
	}
	secSchemes := findMappingValue(comp, "securitySchemes")
	if secSchemes == nil {
		t.Fatal("expected securitySchemes in components")
	}
}
