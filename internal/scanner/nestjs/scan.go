//ff:func feature=scan type=extract control=sequence
//ff:what NestJS 프로젝트를 스캔한다 (미구현 스텁)
package nestjs

import (
	"fmt"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func Scan(root string) (*scanner.ScanResult, error) {
	return nil, fmt.Errorf("nestjs scanner not yet implemented")
}
