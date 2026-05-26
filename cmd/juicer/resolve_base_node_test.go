//ff:func feature=scan type=test control=sequence
//ff:what TestResolveBaseNode 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveBaseNode(t *testing.T) {
	// empty baseFile and no file in root → nil
	dir := t.TempDir()
	if got := resolveBaseNode("", dir); got != nil {
		t.Fatal("expected nil when no base spec found")
	}

	// create a valid openapi.yaml
	spec := `openapi: "3.0.0"
info:
  title: test
  version: "1.0"
paths: {}
`
	specPath := filepath.Join(dir, "openapi.yaml")
	os.WriteFile(specPath, []byte(spec), 0o644)

	// explicit baseFile path
	if got := resolveBaseNode(specPath, dir); got == nil {
		t.Fatal("expected non-nil node for valid spec")
	}

	// auto-discover via empty baseFile
	if got := resolveBaseNode("", dir); got == nil {
		t.Fatal("expected non-nil node via auto-discover")
	}

	// invalid file path
	if got := resolveBaseNode("/nonexistent/file.yaml", dir); got != nil {
		t.Fatal("expected nil for invalid path")
	}
}
