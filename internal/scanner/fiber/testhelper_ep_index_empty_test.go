//ff:func feature=scan type=test control=sequence
//ff:what epIndexEmpty 테스트 헬퍼
package fiber

func epIndexEmpty() map[struct {
	file string
	line int
}]int {
	return map[struct {
		file string
		line int
	}]int{}
}
