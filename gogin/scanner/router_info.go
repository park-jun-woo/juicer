//ff:type feature=scan type=model
//ff:what routerInfo 데이터 구조
package scanner

type routerInfo struct {
	prefix     string
	middleware []string
}
