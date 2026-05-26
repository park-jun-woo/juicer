//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 프로젝트 루트의 Go+Gin 코드에서 엔드포인트를 스캔한다
package gogin

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/packages"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func Scan(root string) (*scanner.ScanResult, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolving path: %w", err)
	}

	cfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedSyntax |
			packages.NeedTypes |
			packages.NeedTypesInfo |
			packages.NeedDeps,
		Dir: absRoot,
	}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, fmt.Errorf("loading packages: %w", err)
	}

	for _, pkg := range pkgs {
		for _, e := range pkg.Errors {
			fmt.Fprintf(os.Stderr, "warning: %v\n", e)
		}
	}

	endpoints, handlerExprsMap := extractRoutes(pkgs, absRoot)
	analyzeHandlers(pkgs, endpoints, absRoot, handlerExprsMap)

	// gin 경로(:param)를 OpenAPI 경로({param})로 변환
	for i := range endpoints {
		endpoints[i].Path = ginPathToOpenAPI(endpoints[i].Path)
	}

	return &scanner.ScanResult{Endpoints: endpoints}, nil
}
