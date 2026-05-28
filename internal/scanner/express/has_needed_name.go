//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 이름 목록에 미해결 스키마가 하나라도 있는지 확인한다
package express

func hasNeededName(names []string, unresolvedSet map[string]bool) bool {
	for _, name := range names {
		if unresolvedSet[name] {
			return true
		}
	}
	return false
}
