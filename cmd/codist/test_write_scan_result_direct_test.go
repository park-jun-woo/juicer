//ff:func feature=scan type=test control=sequence
//ff:what writeScanResult YAML/JSON/OpenAPI 및 파일/stdout 출력 분기 직접 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestWriteScanResultDirect(t *testing.T) {
	res := &scanner.ScanResult{Endpoints: []scanner.Endpoint{{Method: "GET", Path: "/x"}}}
	dir := t.TempDir()

	// YAML to file
	out := filepath.Join(dir, "y.yaml")
	if err := writeScanResult(res, dir, scanOptions{outFile: out}); err != nil {
		t.Fatalf("yaml: %v", err)
	}
	if b, _ := os.ReadFile(out); len(b) == 0 {
		t.Error("yaml file empty")
	}

	// JSON to file
	outj := filepath.Join(dir, "j.json")
	if err := writeScanResult(res, dir, scanOptions{jsonOut: true, outFile: outj}); err != nil {
		t.Fatalf("json: %v", err)
	}

	// OpenAPI to file
	outo := filepath.Join(dir, "o.yaml")
	if err := writeScanResult(res, dir, scanOptions{openapi: true, outFile: outo}); err != nil {
		t.Fatalf("openapi: %v", err)
	}

	// stdout path (no outFile)
	if err := writeScanResult(res, dir, scanOptions{}); err != nil {
		t.Errorf("stdout: %v", err)
	}
}
