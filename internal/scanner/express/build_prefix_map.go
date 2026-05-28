//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 마운트 엔트리 목록에서 파일 경로→prefix 매핑을 생성한다
package express

func buildPrefixMap(mounts []mountEntry) map[string]string {
	prefixMap := make(map[string]string)
	for _, m := range mounts {
		if m.filePath != "" {
			prefixMap[m.filePath] = m.prefix
		}
	}
	propagateParentPrefixes(mounts, prefixMap)
	return prefixMap
}
