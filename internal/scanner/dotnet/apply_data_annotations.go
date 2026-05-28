//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what Data Annotation 어트리뷰트를 Field에 반영한다
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyDataAnnotations(node *sitter.Node, src []byte, f *scanner.Field) {
	if hasAttribute(node, src, AttrRequired) {
		f.Validate = "required"
	}

	if attr := findAttribute(node, src, AttrStringLength); attr != nil {
		args := attributeIntArgs(attr, src)
		if len(args) >= 1 {
			f.MaxLength = intPtr(args[0])
		}
	}

	if attr := findAttribute(node, src, AttrMaxLength); attr != nil {
		args := attributeIntArgs(attr, src)
		if len(args) >= 1 {
			f.MaxLength = intPtr(args[0])
		}
	}

	if attr := findAttribute(node, src, AttrMinLength); attr != nil {
		args := attributeIntArgs(attr, src)
		if len(args) >= 1 {
			f.MinLength = intPtr(args[0])
		}
	}

	if attr := findAttribute(node, src, AttrRange); attr != nil {
		args := attributeIntArgs(attr, src)
		if len(args) >= 2 {
			f.Minimum = intPtr(args[0])
			f.Maximum = intPtr(args[1])
		}
	}

	if hasAttribute(node, src, AttrEmailAddress) {
		f.Validate = appendValidate(f.Validate, "email")
	}
}
