//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what Supabase Edge Functions 프로젝트를 스캔하여 엔드포인트를 추출한다
package supafunc

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
	funcFiles, err := findFunctions(absRoot)
	if err != nil {
		return nil, fmt.Errorf("finding functions: %w", err)
	}
	if len(funcFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}
	var endpoints []scanner.Endpoint
	for _, f := range funcFiles {
		fi, err := parseFile(f)
		if err != nil {
			continue
		}
		eps := extractEndpoint(fi, absRoot)
		endpoints = append(endpoints, eps...)
	}
	return &scanner.ScanResult{Endpoints: endpoints}, nil
}
