//ff:func feature=scan type=extract control=sequence topic=django
//ff:what Django + DRF 프로젝트를 스캔하여 엔드포인트를 추출한다
package django

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// Scan scans a Django + DRF project root and extracts endpoints.
// Pass 1: collect URL routing structure (urlpatterns, include, router.register).
// Pass 2: extract views — ViewSets, APIViews, @api_view functions.
// Pass 3: resolve Serializer fields for request/response schemas.
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
	endpoints := buildAllEndpoints(files)
	return &scanner.ScanResult{Endpoints: endpoints}, nil
}
