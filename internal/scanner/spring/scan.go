//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what Spring Boot 프로젝트를 스캔하여 엔드포인트를 추출한다
package spring

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func Scan(root string) (*scanner.ScanResult, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolving path: %w", err)
	}
	javaFiles, err := findJavaFiles(absRoot)
	if err != nil {
		return nil, fmt.Errorf("finding java files: %w", err)
	}
	if len(javaFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}

	files := parseAllFiles(absRoot, javaFiles)
	controllers := collectControllers(files)
	endpoints, dtoReqs := buildAllEndpoints(controllers, absRoot)
	resolveAllDTOs(dtoReqs, endpoints, absRoot)

	return &scanner.ScanResult{Endpoints: endpoints}, nil
}
