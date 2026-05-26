//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what NestJS 프로젝트를 스캔하여 엔드포인트를 추출한다
package nestjs

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// Scan scans a NestJS project root and extracts endpoints.
// Pass 1: collect routes from @Controller classes.
// Pass 2: resolve DTO types from import paths.
func Scan(root string) (*scanner.ScanResult, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolving path: %w", err)
	}
	tsFiles, err := findTSFiles(absRoot)
	if err != nil {
		return nil, fmt.Errorf("finding ts files: %w", err)
	}
	if len(tsFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}
	globalPrefix := detectGlobalPrefix(absRoot)
	uriVersioning := detectURIVersioning(absRoot)
	allControllers := collectControllers(tsFiles, absRoot)
	endpoints, dtoReqs := buildAllEndpoints(globalPrefix, uriVersioning, allControllers, absRoot)
	resolveAllDTOs(dtoReqs, endpoints)
	return &scanner.ScanResult{Endpoints: endpoints}, nil
}
