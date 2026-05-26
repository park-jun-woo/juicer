//ff:func feature=scan type=convert control=sequence
//ff:what 파라미터 타입 문자열에서 OpenAPI schema map을 생성한다
package scanner

import "strings"

// buildParamSchema creates an OpenAPI schema map from a type string.
// It handles the "type:format" convention (e.g. "string:uuid") and falls back
// to "string" for empty input.
func buildParamSchema(typ string) map[string]any {
	if typ == "" {
		return map[string]any{"type": "string"}
	}
	if i := strings.IndexByte(typ, ':'); i >= 0 {
		return map[string]any{
			"type":   typ[:i],
			"format": typ[i+1:],
		}
	}
	return map[string]any{"type": typ}
}
