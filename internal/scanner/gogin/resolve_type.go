//ff:func feature=scan type=extract control=sequence
//ff:what types.Type에서 named type 이름과 struct 필드를 추출한다
package gogin

import (
	"go/types"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func resolveType(t types.Type) (typeName string, fields []scanner.Field) {
	// 포인터 unwrap
	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}

	// 슬라이스/배열 → 요소 타입에서 추출, 이름에 [] 접두
	isSlice := false
	if sl, ok := t.(*types.Slice); ok {
		t = sl.Elem()
		isSlice = true
		if ptr, ok := t.(*types.Pointer); ok {
			t = ptr.Elem()
		}
	} else if arr, ok := t.(*types.Array); ok {
		t = arr.Elem()
		isSlice = true
		if ptr, ok := t.(*types.Pointer); ok {
			t = ptr.Elem()
		}
	}

	// named type에서 이름 추출
	if named, ok := t.(*types.Named); ok {
		if wk, ok := wellKnownType(named); ok {
			return slicePrefix(isSlice) + wk, nil
		}
		typeName = named.Obj().Name()
		t = named.Underlying()
	}

	if isSlice && typeName != "" {
		typeName = "[]" + typeName
	}

	st, ok := t.(*types.Struct)
	if !ok {
		return typeName, nil
	}

	visited := make(map[string]bool)
	fields = extractFields(st, visited)
	return typeName, fields
}

