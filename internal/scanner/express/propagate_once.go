//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 마운트 엔트리를 한 번 순회하며 부모 prefix를 전이하고 변경 여부를 반환한다
package express

func propagateOnce(mounts []mountEntry, prefixMap map[string]string) bool {
	changed := false
	for _, m := range mounts {
		combined := applyParentPrefix(prefixMap, m)
		if combined != "" && prefixMap[m.filePath] != combined {
			prefixMap[m.filePath] = combined
			changed = true
		}
	}
	return changed
}
