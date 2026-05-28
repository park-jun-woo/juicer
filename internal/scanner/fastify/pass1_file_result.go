//ff:type feature=scan type=model topic=fastify
//ff:what Pass 1 단일 파일 처리 결과 구조체
package fastify

type pass1FileResult struct {
	fi        *fileInfo
	instances map[string]bool
	mounts    []pluginMount
}
