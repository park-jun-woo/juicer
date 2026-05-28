//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what Flask 프로젝트를 스캔하여 엔드포인트를 추출한다
package flask

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// Scan scans a Flask project root and extracts endpoints.
// Pass 1: collect Blueprint structure (Blueprint instances, register_blueprint chains).
// Pass 2: extract routes from decorated handler functions.
func Scan(root string) (*scanner.ScanResult, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolving path: %w", err)
	}
	pyFiles, err := findPyFiles(absRoot)
	if err != nil {
		return nil, fmt.Errorf("finding py files: %w", err)
	}
	if len(pyFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}

	files := parseAllFiles(absRoot, pyFiles)
	if len(files) == 0 {
		return &scanner.ScanResult{}, nil
	}

	bpPrefixes := resolveBlueprintPrefixes(files)

	var endpoints []scanner.Endpoint
	for _, fi := range files {
		routes := extractRoutes(fi.root, fi.src, bpPrefixes, fi.relPath)
		for _, ri := range routes {
			endpoints = append(endpoints, buildEndpoint(ri))
		}
	}

	return &scanner.ScanResult{Endpoints: endpoints}, nil
}
