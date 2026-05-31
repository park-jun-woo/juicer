//ff:type feature=scan type=model topic=nestjs
//ff:what 중첩 DTO/enum 스키마 등록 레지스트리 구조체
package nestjs

import "github.com/park-jun-woo/codistill/internal/scanner"

// schemaRegistry drives the recursive registration of nested DTO/enum named
// types into the component schemas map during a BFS over field references.
type schemaRegistry struct {
	cache     map[string][]scanner.Field
	schemas   map[string]any
	queue     []schemaJob
	processed map[string]bool
	topLevel  map[string]bool
}
