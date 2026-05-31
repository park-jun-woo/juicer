//ff:func feature=scan type=convert control=sequence topic=nestjs
//ff:what named DTO/enum 타입이면 scanner.Field에 $ref(및 array)를 설정한다
package nestjs

import "github.com/park-jun-woo/codistill/internal/scanner"

// applyNamedRef sets sf.Ref (and Type=array for array forms) when tsType names
// a non-builtin DTO/enum. Inlined enums (already carrying Enum values) are left
// untouched so they remain inline string enums.
func applyNamedRef(sf *scanner.Field, tsType string) {
	if len(sf.Enum) > 0 {
		return
	}
	base, isArr, ok := namedDTOType(tsType)
	if !ok {
		return
	}
	sf.Ref = base
	if isArr {
		sf.Type = "array"
	}
}
