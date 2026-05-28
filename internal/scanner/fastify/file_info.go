//ff:type feature=scan type=model topic=fastify
//ff:what 파싱된 TypeScript 파일 정보 구조체
package fastify

import sitter "github.com/smacker/go-tree-sitter"

type fileInfo struct {
	Path string
	Root *sitter.Node
	Src  []byte
}
