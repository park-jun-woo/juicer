//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 플러그인 마운트의 파일 경로를 import 매핑 또는 직접 경로에서 해석한다
package fastify

func resolvePluginFilePaths(mounts []pluginMount, imports map[string]string, absRoot string) {
	for i := range mounts {
		ref := mounts[i].PluginRef
		if ref == "" || ref == inlineRef {
			continue
		}
		if resolved, ok := imports[ref]; ok {
			mounts[i].FilePath = resolved
			continue
		}
		resolved := resolveRelativePath(absRoot, ref)
		if resolved != "" {
			mounts[i].FilePath = resolved
		}
	}
}
