//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what declOfType 테스트 헬퍼
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func declOfType(t *testing.T, src string) ([]*sitter.Node, *fileInfo) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	return findAllByType(fi.Root, "lexical_declaration"), fi
}
