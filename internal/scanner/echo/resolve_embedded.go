//ff:func feature=scan type=extract control=sequence
//ff:what 임베딩된 타입의 필드를 재귀적으로 추출한다
package echo

import (
	"go/types"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveEmbedded(t types.Type, visited map[string]bool) []scanner.Field {
	t = unwrapPointer(t)

	key := t.String()
	if visited[key] {
		return nil
	}
	visited[key] = true
	defer delete(visited, key)

	if named, ok := t.(*types.Named); ok {
		t = named.Underlying()
	}

	st, ok := t.(*types.Struct)
	if !ok {
		return nil
	}

	return extractFields(st, visited)
}
