//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 모든 파일에서 absPath → (varName → prefix) 글로벌 맵을 생성한다
package fastapi

// buildGlobalPrefixMap creates a map from file absPath to (varName -> prefix).
func buildGlobalPrefixMap(files []fileInfo) map[string]map[string]string {
	result := make(map[string]map[string]string)
	for _, fi := range files {
		if len(fi.prefixes) == 0 {
			continue
		}
		m := make(map[string]string, len(fi.prefixes))
		for k, v := range fi.prefixes {
			m[k] = v
		}
		result[fi.absPath] = m
	}
	return result
}
