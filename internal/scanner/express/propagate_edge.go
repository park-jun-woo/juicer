//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 부모 prefix들을 한 엣지의 세그먼트와 결합해 자식에 추가하고 변경 여부를 반환한다
package express

func propagateEdge(prefixes map[routerKey][]string, e routerEdge, pps []string) bool {
	changed := false
	for _, pp := range pps {
		if appendUniquePrefix(prefixes, e.child, joinExpressPath(pp, e.seg)) {
			changed = true
		}
	}
	return changed
}
