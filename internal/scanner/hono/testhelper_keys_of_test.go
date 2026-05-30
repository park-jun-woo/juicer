//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what keysOf 테스트 헬퍼
package hono

func keysOf(m map[string][]byte) []string {
	var k []string
	for key := range m {
		k = append(k, key)
	}
	return k
}
