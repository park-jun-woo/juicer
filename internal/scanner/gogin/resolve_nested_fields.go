//ff:func feature=scan type=extract control=sequence
//ff:what 필드 타입이 struct이면 중첩 필드를 추출한다
package gogin

import (
	"go/types"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveNestedFields(t types.Type, visited map[string]bool) []scanner.Field {
	t = unwrapPointer(t)

	// 슬라이스/배열 — 요소 타입에서 struct 추출
	if sl, ok := t.(*types.Slice); ok {
		t = unwrapPointer(sl.Elem())
	}
	if arr, ok := t.(*types.Array); ok {
		t = unwrapPointer(arr.Elem())
	}

	key := t.String()
	if visited[key] {
		return nil
	}
	visited[key] = true
	defer delete(visited, key)

	if named, ok := t.(*types.Named); ok {
		if _, ok := wellKnownType(named); ok {
			return nil // time.Time 등은 struct 전개하지 않음
		}
		t = named.Underlying()
	}

	st, ok := t.(*types.Struct)
	if !ok {
		return nil
	}

	return extractFields(st, visited)
}

