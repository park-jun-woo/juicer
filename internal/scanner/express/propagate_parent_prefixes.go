//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 수렴 루프로 상위 prefix를 하위 파일에 전이한다
package express

func propagateParentPrefixes(mounts []mountEntry, prefixMap map[string]string) {
	const maxIter = 10
	for i := 0; i < maxIter; i++ {
		if !propagateOnce(mounts, prefixMap) {
			break
		}
	}
}
