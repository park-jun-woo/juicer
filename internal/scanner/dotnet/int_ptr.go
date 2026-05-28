//ff:func feature=scan type=convert control=sequence topic=dotnet
//ff:what int 값의 포인터를 반환한다
package dotnet

func intPtr(v int) *int {
	return &v
}
