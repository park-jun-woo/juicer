//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what JSON Schema 숫자 제약 (minimum, maximum, minLength, maxLength)을 필드에 적용한다
package fastify

import (
	"strconv"

	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func applyNumericConstraints(f *scanner.Field, propNode *sitter.Node, src []byte) {
	pairs := []struct {
		key  string
		dest **int
	}{
		{"minimum", &f.Minimum},
		{"maximum", &f.Maximum},
		{"minLength", &f.MinLength},
		{"maxLength", &f.MaxLength},
	}
	for _, p := range pairs {
		s := extractPairStringOrIdent(propNode, src, p.key)
		if s == "" {
			continue
		}
		if v, err := strconv.Atoi(s); err == nil {
			*p.dest = &v
		}
	}
}
