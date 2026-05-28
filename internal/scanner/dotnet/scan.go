//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what ASP.NET Core 프로젝트를 스캔하여 엔드포인트를 추출한다
package dotnet

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
	csFiles, err := findCsFiles(absRoot)
	if err != nil {
		return nil, fmt.Errorf("finding cs files: %w", err)
	}
	if len(csFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}

	files := parseAllFiles(absRoot, csFiles)
	controllers := collectControllers(files)
	endpoints, dtoReqs := buildAllEndpoints(controllers, absRoot)

	groups := extractMapGroups(files)
	minimalEndpoints := extractMinimalAPIs(files, groups)
	endpoints = append(endpoints, minimalEndpoints...)

	resolveAllDTOs(dtoReqs, endpoints, files)

	return &scanner.ScanResult{Endpoints: endpoints}, nil
}
