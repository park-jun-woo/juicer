//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what keysOf — map 키 목록 추출 테스트 헬퍼
package nestjs

func keysOf(m map[string]any) []string {
	var k []string
	for key := range m {
		k = append(k, key)
	}
	return k
}
