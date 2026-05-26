//ff:func feature=scan type=extract control=sequence
//ff:what TestMergeSpec_NoOverlap 테스트
package scanner

import (
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMergeSpec_NoOverlap(t *testing.T) {
	// scan에만 /users가 있고, base에는 paths가 비어있음
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
paths: {}
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

	// base의 info가 보존되어야 한다
	if !strings.Contains(s, "My API") {
		t.Error("expected base info title preserved")
	}
	// scan의 /users가 포함되어야 한다
	if !strings.Contains(s, "/users") {
		t.Error("expected /users path in merged result")
	}
}
