//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 타입명으로 struct 정의를 찾아 필드를 반환한다(캐시 사용)
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveStructFields(typeName string, idx structIndex, cache map[string][]scanner.Field) []scanner.Field {
	if cached, ok := cache[typeName]; ok {
		return cached
	}
	entry, ok := idx[typeName]
	if !ok {
		return nil
	}

	for _, structNode := range findAllByType(entry.file.root, "struct_item") {
		if !structNameMatches(structNode, entry.file.src, typeName) {
			continue
		}
		fields := extractStructFields(structNode, entry.file.src)
		cache[typeName] = fields
		return fields
	}

	return nil
}
