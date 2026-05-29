//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 단일 파일의 autoload register 호출들을 디렉터리 워크 기반 마운트로 변환한다
package fastify

import "path/filepath"

func fileAutoloadMounts(path string, fi *fileInfo, absRoot string) []pluginMount {
	names := autoloadVarNames(fi)
	if len(names) == 0 {
		return nil
	}
	srcDir := filepath.Dir(path)
	var mounts []pluginMount
	for _, call := range findAllByType(fi.Root, "call_expression") {
		dir, prefix, ok := extractAutoloadCall(call, fi.Src, names)
		if !ok {
			continue
		}
		baseDir := filepath.Join(srcDir, dir)
		mounts = append(mounts, autoloadDirMounts(path, baseDir, prefix)...)
	}
	return mounts
}
