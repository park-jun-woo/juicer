//ff:func feature=scan type=extract control=sequence
//ff:what TestMergeSpec_FullOverlap 테스트
package scanner

import (
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMergeSpec_FullOverlap(t *testing.T) {
	// scan과 base 모두 /users GET을 가짐
	scanResult := &ScanResult{
		Endpoints: []Endpoint{
			{Method: "GET", Path: "/users", Handler: "h.ListUsers"},
		},
	}
	scanNode := buildSpecNode(scanResult)

	baseYAML := `openapi: "3.0.3"
info:
  title: "My API"
  version: "1.0.0"
paths:
  /users:
    get:
      summary: "List all users"
      description: "Returns a list of users"
`
	var doc yaml.Node
	yaml.Unmarshal([]byte(baseYAML), &doc)
	baseNode := doc.Content[0]

	merged := mergeSpec(scanNode, baseNode, scanResult)

	out, err := yaml.Marshal(merged)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	s := string(out)

	// base의 description이 보존되어야 한다
	if !strings.Contains(s, "Returns a list of users") {
		t.Error("expected base description preserved in merged result")
	}
	// base의 summary가 보존되어야 한다
	if !strings.Contains(s, "List all users") {
		t.Error("expected base summary preserved in merged result")
	}
}
