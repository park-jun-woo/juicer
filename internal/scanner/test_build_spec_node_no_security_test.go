//ff:func feature=scan type=test control=sequence
//ff:what TestBuildSpecNode_NoSecurity 테스트
package scanner

import "testing"

func TestBuildSpecNode_NoSecurity(t *testing.T) {
	// endpoints without auth should not produce securitySchemes
	result := &ScanResult{
		Endpoints: []Endpoint{
			{
				Method: "GET",
				Path:   "/api/health",
			},
		},
	}
	node := buildSpecNode(result)
	comp := findMappingValue(node, "components")
	if comp != nil {
		secSchemes := findMappingValue(comp, "securitySchemes")
		if secSchemes != nil {
			t.Fatal("expected no securitySchemes when no auth endpoints")
		}
	}
}
