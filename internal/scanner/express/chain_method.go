//ff:type feature=scan type=model topic=express
//ff:what 체인 메서드 정보 구조체
package express

type chainMethod struct {
	method     string
	handler    string
	middleware []string
	line       int
}
