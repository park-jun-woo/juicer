//ff:func feature=scan type=extract control=sequence topic=express
//ff:what Express 프로젝트를 스캔하여 엔드포인트를 추출한다 (2-pass)
package express

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
	tsFiles, err := findTSFiles(absRoot)
	if err != nil {
		return nil, fmt.Errorf("finding ts files: %w", err)
	}
	if len(tsFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}
	ctx := scanPass1(tsFiles, absRoot)
	endpoints := scanPass2(ctx, absRoot)
	return &scanner.ScanResult{Endpoints: endpoints}, nil
}
