//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what autoload 디렉터리 하위의 각 .ts 파일을 디렉터리 구조 기반 prefix로 마운트한다
package fastify

func autoloadDirMounts(srcFile, baseDir, basePrefix string) []pluginMount {
	files := collectDirTSFiles(baseDir)
	var mounts []pluginMount
	for _, path := range files {
		prefix, ok := autoloadFilePrefix(baseDir, path, basePrefix)
		if !ok {
			continue
		}
		mounts = append(mounts, pluginMount{PluginRef: path, Prefix: prefix, FilePath: path, SourceFile: srcFile})
	}
	return mounts
}
