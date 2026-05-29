//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what 서브앱 정의 파일의 Hono 변수(cross-file는 전부, 동일 파일은 해당 변수)에 prefix를 기록한다
package hono

func assignSubAppPrefix(defFile, srcFile, subAppName, fullPrefix string, prefixMap map[string]string, honoVars map[string]map[string]bool) bool {
	if defFile == srcFile {
		key := prefixKey(srcFile, subAppName)
		if prefixMap[key] == fullPrefix {
			return false
		}
		prefixMap[key] = fullPrefix
		return true
	}
	changed := false
	for varName := range honoVars[defFile] {
		key := prefixKey(defFile, varName)
		if prefixMap[key] != fullPrefix {
			prefixMap[key] = fullPrefix
			changed = true
		}
	}
	return changed
}
