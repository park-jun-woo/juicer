//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 인라인 라우터 prefix 맵을 한 번 순회하며 부모 prefix를 합산하고 변경 여부를 반환한다
package express

func convergeLocalPrefixesOnce(entries []mountEntry, localPrefixes map[string]string) bool {
	changed := false
	for _, e := range entries {
		if e.filePath != "" || e.varName == "" {
			continue
		}
		parentPrefix, ok := localPrefixes[e.sourceRouter]
		if !ok {
			continue
		}
		combined := joinExpressPath(parentPrefix, e.prefix)
		if localPrefixes[e.varName] != combined {
			localPrefixes[e.varName] = combined
			changed = true
		}
	}
	return changed
}
