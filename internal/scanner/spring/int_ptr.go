//ff:func feature=scan type=convert control=sequence topic=spring
//ff:what int 값의 포인터를 반환한다
package spring

func intPtr(v int) *int {
	return &v
}
