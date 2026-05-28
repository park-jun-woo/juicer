//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 함수 파라미터의 타입 어노테이션에서 express.Router 타입을 찾아 등록한다
package express

func collectParamRoutersFromType(fi *fileInfo, routers map[string]bool) {
	funcTypes := []string{"arrow_function", "function_declaration"}
	for _, ft := range funcTypes {
		for _, fn := range findAllByType(fi.Root, ft) {
			collectRouterParamsFromFunc(fn, fi.Src, routers)
		}
	}
}
