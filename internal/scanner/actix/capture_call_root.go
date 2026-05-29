//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what scoped_identifier 노드가 web::scope/web::resource면 결과에 처음 한 번 기록한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func captureCallRoot(n *sitter.Node, src []byte, result *string) {
	if *result != "" || n.Type() != "scoped_identifier" {
		return
	}
	text := nodeText(n, src)
	if text == "web::scope" || text == "web::resource" {
		*result = text
	}
}
