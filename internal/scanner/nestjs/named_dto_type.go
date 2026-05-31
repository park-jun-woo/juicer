//ff:func feature=scan type=convert control=sequence topic=nestjs
//ff:what TS 타입 문자열에서 named DTO/enum 기본 타입명을 추출한다
package nestjs

import "strings"

// namedDTOType returns the named (non-builtin) base type referenced by a TS
// type string and whether it is an array. It returns ("", false, false) for
// primitives, builtins, and unparseable/union types — those should not become
// component $refs. e.g. "AlbumUserResponseDto[]" -> ("AlbumUserResponseDto",
// true, true); "string" -> ("", false, false).
func namedDTOType(ts string) (baseName string, isArray bool, ok bool) {
	ts = strings.TrimSpace(unwrapPromise(strings.TrimSpace(ts)))
	ts, isArray = stripArrayWrapper(ts)
	if !isSimpleTypeName(ts) || isBuiltinTSType(ts) {
		return "", false, false
	}
	return ts, isArray, true
}
