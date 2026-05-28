//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what Fastify() 인스턴스 변수와 함수 파라미터의 fastify 바인딩을 수집한다
package fastify

func collectInstances(fi *fileInfo) map[string]bool {
	instances := make(map[string]bool)
	collectFastifyVars(fi, instances)
	funcTypes := []string{"arrow_function", "function_declaration", "function", "function_expression"}
	for _, ft := range funcTypes {
		for _, fn := range findAllByType(fi.Root, ft) {
			collectFuncParamInstance(fn, fi.Src, instances)
		}
	}
	return instances
}
