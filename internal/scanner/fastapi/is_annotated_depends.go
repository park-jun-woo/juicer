//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 타입명이 Annotated Depends 별칭 또는 인라인 패턴인지 판별한다
package fastapi

import "strings"

// isAnnotatedDepends checks if the typeName resolves to a Depends-based type,
// either via the alias map or via inline Annotated[..., Depends(...)].
func isAnnotatedDepends(typeName string, aliasMap map[string]string) bool {
	if aliasMap != nil {
		if _, ok := aliasMap[typeName]; ok {
			return true
		}
	}
	return strings.HasPrefix(typeName, "Annotated[") && strings.Contains(typeName, "Depends(")
}
