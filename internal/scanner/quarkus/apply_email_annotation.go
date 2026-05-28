//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what @Email 어노테이션으로 타입을 string:email로 설정한다
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyEmailAnnotation(field *sitter.Node, src []byte, f *scanner.Field) {
	if hasAnnotation(field, src, AnnEmail) {
		f.Type = "string:email"
	}
}
