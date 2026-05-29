//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 모든 파일의 @fastify/autoload register 호출에서 디렉터리 기반 플러그인 마운트를 생성한다
package fastify

func collectAutoloadMounts(parsed map[string]*fileInfo, absRoot string) []pluginMount {
	var mounts []pluginMount
	for path, fi := range parsed {
		mounts = append(mounts, fileAutoloadMounts(path, fi, absRoot)...)
	}
	return mounts
}
