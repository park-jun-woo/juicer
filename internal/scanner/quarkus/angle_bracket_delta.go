//ff:func feature=scan type=extract control=selection topic=quarkus
//ff:what 꺾쇠괄호의 깊이 변화량을 반환한다
package quarkus

func angleBracketDelta(ch rune) int {
	switch ch {
	case '<':
		return 1
	case '>':
		return -1
	default:
		return 0
	}
}
