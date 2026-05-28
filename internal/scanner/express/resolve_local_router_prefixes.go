//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 인라인 라우터의 prefix를 파일 마운트 엔트리에 합산한다
package express

func resolveLocalRouterPrefixes(entries []mountEntry) []mountEntry {
	// Phase A: build and converge localPrefixes
	localPrefixes := make(map[string]string)
	for _, e := range entries {
		if e.filePath == "" && e.varName != "" {
			localPrefixes[e.varName] = e.prefix
		}
	}
	if len(localPrefixes) == 0 {
		return entries
	}
	for i := 0; i < 5; i++ {
		if !convergeLocalPrefixesOnce(entries, localPrefixes) {
			break
		}
	}

	// Phase B: apply resolved prefixes to file-mount entries (one pass)
	var result []mountEntry
	for _, e := range entries {
		if e.filePath == "" {
			continue
		}
		parentPrefix, ok := localPrefixes[e.sourceRouter]
		if ok {
			e.prefix = joinExpressPath(parentPrefix, e.prefix)
		}
		result = append(result, e)
	}
	return result
}
