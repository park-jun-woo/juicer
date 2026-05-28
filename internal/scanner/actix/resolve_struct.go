//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 타입명으로 struct 정의를 찾아 필드를 반환한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

type structIndex map[string]*structEntry

type structEntry struct {
	file       *fileInfo
	structName string
}

func buildStructIndex(files []*fileInfo) structIndex {
	idx := make(structIndex)
	for _, fi := range files {
		for _, structNode := range findAllByType(fi.root, "struct_item") {
			nameNode := findChildByType(structNode, "type_identifier")
			if nameNode == nil {
				continue
			}
			name := nodeText(nameNode, fi.src)
			idx[name] = &structEntry{
				file:       fi,
				structName: name,
			}
		}
	}
	return idx
}

func resolveStructFields(typeName string, idx structIndex, cache map[string][]scanner.Field) []scanner.Field {
	if cached, ok := cache[typeName]; ok {
		return cached
	}

	entry, ok := idx[typeName]
	if !ok {
		return nil
	}

	for _, structNode := range findAllByType(entry.file.root, "struct_item") {
		nameNode := findChildByType(structNode, "type_identifier")
		if nameNode == nil {
			continue
		}
		name := nodeText(nameNode, entry.file.src)
		if name != typeName {
			continue
		}
		fields := extractStructFields(structNode, entry.file.src)
		cache[typeName] = fields
		return fields
	}

	return nil
}
