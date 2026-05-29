//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what 단일 route 그룹의 prefix를 부모에서 합성하여 서브앱(같은 파일 또는 import된 정의 파일)에 부여한다
package hono

func applyGroupPrefix(g routeGroup, prefixMap map[string]string, honoVars map[string]map[string]bool, imports map[string]map[string]string) bool {
	parentPrefix := prefixMap[prefixKey(g.SourceFile, g.ParentVar)]
	fullPrefix := joinHonoPath(parentPrefix, g.Prefix)
	defFile := g.SourceFile
	if fileImports, ok := imports[g.SourceFile]; ok {
		if resolved, ok := fileImports[g.SubAppName]; ok {
			defFile = resolved
		}
	}
	return assignSubAppPrefix(defFile, g.SourceFile, g.SubAppName, fullPrefix, prefixMap, honoVars)
}
