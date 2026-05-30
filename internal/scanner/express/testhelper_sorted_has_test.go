//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what sortedHas 테스트 헬퍼
package express

func sortedHas(names []string, want string) bool {
	for _, n := range names {
		if n == want {
			return true
		}
	}
	return false
}
