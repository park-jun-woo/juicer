//ff:type feature=scan type=model topic=express
//ff:what 배열 object에서 추출한 path/route 쌍 구조체
package express

type pathRouteEntry struct {
	path     string
	routeVar string
}
