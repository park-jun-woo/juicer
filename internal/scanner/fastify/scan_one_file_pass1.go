//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 단일 파일의 Pass 1: 파싱, 인스턴스 수집, import 해석, 플러그인 마운트 수집
package fastify

func scanOneFilePass1(path, absRoot string) *pass1FileResult {
	fi, err := parseFile(path)
	if err != nil {
		return nil
	}
	instances := collectInstances(fi)
	imports := resolveImports(fi, absRoot)
	mounts := collectPlugins(fi, instances)
	resolvePluginFilePaths(mounts, imports, absRoot)
	return &pass1FileResult{fi: fi, instances: instances, mounts: mounts}
}
