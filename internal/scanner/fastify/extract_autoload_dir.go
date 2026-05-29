//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what autoload 옵션 객체의 dir 값에서 디렉터리 상대 경로를 추출한다
package fastify

import (
	"path/filepath"

	sitter "github.com/smacker/go-tree-sitter"
)

func extractAutoloadDir(opts *sitter.Node, src []byte) string {
	val := findPairValue(opts, src, "dir")
	if val == nil {
		return ""
	}
	if val.Type() == "string" || val.Type() == "template_string" {
		return unquoteTS(nodeText(val, src))
	}
	segs := dirStringSegments(val, src)
	if len(segs) == 0 {
		return ""
	}
	return filepath.Join(segs...)
}
