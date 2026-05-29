//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what TypeBox 스칼라의 옵션 객체(format/제약/enum)를 필드에 적용한다
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func applyTypeBoxOptions(f *scanner.Field, call *sitter.Node, src []byte) {
	opts := typeBoxFirstArg(call)
	if opts == nil || opts.Type() != "object" {
		return
	}
	applyFormat(f, opts, src)
	applyNumericConstraints(f, opts, src)
	applyEnum(f, opts, src)
}
