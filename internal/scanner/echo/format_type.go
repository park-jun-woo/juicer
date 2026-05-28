//ff:func feature=scan type=extract control=selection
//ff:what types.Type을 간결한 문자열로 변환한다
package echo

import (
	"go/types"
)

func formatType(t types.Type) string {
	switch tt := t.(type) {
	case *types.Basic:
		return tt.Name()
	case *types.Pointer:
		return "*" + formatType(tt.Elem())
	case *types.Slice:
		return "[]" + formatType(tt.Elem())
	case *types.Array:
		return "[]" + formatType(tt.Elem())
	case *types.Map:
		return "map[" + formatType(tt.Key()) + "]" + formatType(tt.Elem())
	case *types.Named:
		obj := tt.Obj()
		pkg := obj.Pkg()
		if pkg == nil {
			return obj.Name()
		}
		return pkg.Name() + "." + obj.Name()
	case *types.Interface:
		return "any"
	case *types.Struct:
		return "object"
	}
	return t.String()
}
