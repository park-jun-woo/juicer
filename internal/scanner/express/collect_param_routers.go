//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 함수 파라미터에서 express.Router 타입 어노테이션을 가진 변수명을 수집한다
package express

func collectParamRouters(fi *fileInfo) map[string]bool {
	routers := make(map[string]bool)
	collectParamRoutersFromType(fi, routers)
	if len(routers) == 0 {
		collectParamRoutersFromUsage(fi, routers)
	}
	return routers
}
