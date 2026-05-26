//ff:func feature=scan type=extract control=sequence
//ff:what TestMergeSpec_DeadSpec 테스트
package scanner

import (
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMergeSpec_DeadSpec(t *testing.T) {
	// scan에는 /users만, base에는 /users와 /legacy가 있음
	// /legacy는 dead spec이므로 제거되어야 한다
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
      summary: "List users"
  /legacy:
    get:
      summary: "Legacy endpoint"
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

	// /users는 유지
	if !strings.Contains(s, "/users") {
		t.Error("expected /users path preserved")
	}
	// /legacy는 제거
	if strings.Contains(s, "/legacy") {
		t.Error("expected /legacy to be dropped as dead spec")
	}
}
