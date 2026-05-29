//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what field_expression이 대상 메서드 호출이면 그 arguments에 fn을 적용한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func invokeIfMethodCall(n, fe *sitter.Node, src []byte, method string, fn func(*sitter.Node)) {
	fid := findChildByType(fe, "field_identifier")
	if fid == nil || nodeText(fid, src) != method {
		return
	}
	if args := findChildByType(n, "arguments"); args != nil {
		fn(args)
	}
}
