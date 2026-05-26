//ff:func feature=scan type=extract control=sequence
//ff:what TestLoadBaseSpec_Valid 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestLoadBaseSpec_Valid(t *testing.T) {
	dir := t.TempDir()
	specPath := filepath.Join(dir, "openapi.yaml")
	content := `openapi: "3.0.3"
info:
  title: "Test API"
  version: "1.0.0"
paths:
  /users:
    get:
      summary: "List users"
`
	os.WriteFile(specPath, []byte(content), 0o644)

	node, err := LoadBaseSpec(specPath)
	if err != nil {
		t.Fatalf("LoadBaseSpec() error: %v", err)
	}
	if node == nil {
		t.Fatal("expected non-nil node")
	}
	if node.Kind != yaml.MappingNode {
		t.Fatalf("expected MappingNode, got %d", node.Kind)
	}
}
