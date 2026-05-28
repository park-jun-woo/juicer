//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what fastify.register() 호출에서 플러그인 + prefix를 수집한다
package fastify

func collectPlugins(fi *fileInfo, instances map[string]bool) []pluginMount {
	var mounts []pluginMount
	for _, call := range findAllByType(fi.Root, "call_expression") {
		pm := extractRegisterCall(call, fi.Src, instances)
		if pm != nil {
			pm.SourceFile = fi.Path
			mounts = append(mounts, *pm)
		}
	}
	return mounts
}
