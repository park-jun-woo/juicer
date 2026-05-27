//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what FastAPI 프로젝트를 스캔하여 엔드포인트를 추출한다
package fastapi

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// Scan scans a FastAPI project root and extracts endpoints.
// Pass 1: collect router structure (FastAPI/APIRouter instances, include_router chains).
// Pass 2: extract routes + params from decorated handler functions.
// Pass 3: resolve Pydantic model types to fill body/response fields.
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
	endpoints, modelReqs := collectEndpoints(files)
	resolveAllModels(modelReqs, endpoints, files)

	return &scanner.ScanResult{Endpoints: endpoints}, nil
}
