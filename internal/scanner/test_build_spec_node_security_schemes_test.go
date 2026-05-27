//ff:func feature=scan type=test control=sequence
//ff:what TestBuildSpecNode_SecuritySchemes 테스트
package scanner

import "testing"

func TestBuildSpecNode_SecuritySchemes(t *testing.T) {
	// endpoints with auth should produce securitySchemes in components
	result := &ScanResult{
		Endpoints: []Endpoint{
			{
				Method:    "GET",
				Path:      "/api/users",
				AuthLevel: "auth_required",
			},
			{
				Method:    "GET",
				Path:      "/api/health",
				AuthLevel: "public",
			},
		},
	}
	node := buildSpecNode(result)
	if node == nil {
		t.Fatal("expected non-nil")
	}

	comp := findMappingValue(node, "components")
	if comp == nil {
		t.Fatal("expected components node")
	}
	secSchemes := findMappingValue(comp, "securitySchemes")
	if secSchemes == nil {
		t.Fatal("expected securitySchemes in components")
	}

	bearerAuth := findMappingValue(secSchemes, "bearerAuth")
	if bearerAuth == nil {
		t.Fatal("expected bearerAuth in securitySchemes")
	}

	// verify operation-level security
	paths := findMappingValue(node, "paths")
	if paths == nil {
		t.Fatal("expected paths")
	}

	// /api/users should have security
	usersPath := findMappingValue(paths, "/api/users")
	if usersPath == nil {
		t.Fatal("expected /api/users path")
	}
	getOp := findMappingValue(usersPath, "get")
	if getOp == nil {
		t.Fatal("expected get operation on /api/users")
	}
	// getOp is a toYAMLNode result, check for security key
	secNode := findMappingValue(getOp, "security")
	if secNode == nil {
		t.Fatal("expected security on /api/users GET")
	}

	// /api/health should NOT have security
	healthPath := findMappingValue(paths, "/api/health")
	if healthPath == nil {
		t.Fatal("expected /api/health path")
	}
	getHealth := findMappingValue(healthPath, "get")
	if getHealth == nil {
		t.Fatal("expected get operation on /api/health")
	}
	if findMappingValue(getHealth, "security") != nil {
		t.Fatal("expected no security on /api/health GET")
	}
}
